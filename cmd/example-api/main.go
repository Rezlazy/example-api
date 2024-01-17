package main

import (
	"context"
	"crypto/rsa"
	server_api "example-api/api/server"
	"example-api/internal/api/http/handler"
	"example-api/internal/api/http/middleware"
	"example-api/internal/api/http/server"
	"example-api/internal/config"
	item_dao "example-api/internal/database/dao/item"
	list_dao "example-api/internal/database/dao/list"
	"example-api/pkg/httputil"
	"example-api/pkg/logger"
	"example-api/pkg/pg"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"net/http"
	"net/url"
)

func main() {
	ctx := context.Background()

	cfg, err := config.Parse()
	if err != nil {
		err = fmt.Errorf("parse config: %w", err)
		logger.Fatal(ctx, err.Error())
	}

	pgxConfig, _ := pgxpool.ParseConfig("")
	pgxConfig.ConnConfig.Host = cfg.PG.Host
	pgxConfig.ConnConfig.Port = cfg.PG.Port
	pgxConfig.ConnConfig.Database = cfg.PG.Database
	pgxConfig.ConnConfig.User = cfg.PG.User
	pgxConfig.ConnConfig.Password = cfg.PG.Password

	pgxPool, err := pgxpool.NewWithConfig(ctx, pgxConfig)
	if err != nil {
		err = fmt.Errorf("create pgx pool: %w", err)
		logger.Fatal(ctx, err.Error())
	}

	pgPGXPool := pg.NewWithPGXPool(pgxPool)

	listDao := list_dao.New()
	itemDao := item_dao.New()

	privateKey := (*rsa.PrivateKey)(cfg.PrivateKey)
	publicKey := (*rsa.PublicKey)(cfg.PublicKey)

	handlerAPI := handler.New(cfg.Auth.Email, cfg.Auth.Password, privateKey, pgPGXPool, listDao, itemDao)
	securityHandler := handler.NewSecurityHandler(publicKey)

	var serverAPI http.Handler
	serverAPI, err = server.NewServer(handlerAPI, securityHandler)
	if err != nil {
		err = fmt.Errorf("create server api: %w", err)
		logger.Fatal(ctx, err.Error())
	}

	serverAPI = middleware.CORSMW(serverAPI)
	serverAPI = httputil.NewRequestTimeout(serverAPI, cfg.Server.DefaultTimeout, nil)

	endpointWithoutHost := fmt.Sprintf(":%d", cfg.Server.Port)
	endpoint := cfg.Server.Host + endpointWithoutHost
	endpointURL := &url.URL{
		Host: endpoint,
	}

	go runSwagger(ctx, endpointURL)

	if err := http.ListenAndServe(endpointWithoutHost, serverAPI); err != nil {
		err = fmt.Errorf("listen and serve: %w", err)
		logger.Fatal(ctx, err.Error())
	}
}

func runSwagger(ctx context.Context, endpoint *url.URL) {
	swaggerHandler, err := httputil.SwaggerHandler(ctx, server_api.SchemaDir, server_api.SchemaFilePath, endpoint)
	if err != nil {
		logger.Error(ctx, err.Error())
	}

	if err := http.ListenAndServe(fmt.Sprintf(":%d", 8081), swaggerHandler); err != nil {
		err = fmt.Errorf("listen and serve: %w", err)
		logger.Error(ctx, err.Error())
	}
}
