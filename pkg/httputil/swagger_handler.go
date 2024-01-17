package httputil

import (
	"context"
	"embed"
	"fmt"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
	"net/url"
)

const (
	schemaDirRoute      = "/schema"
	fullSchemaFileRoute = "/schema.json"
)

func SwaggerHandler(ctx context.Context, schemaFS embed.FS, schemaFilePath string, endpoint string) (http.Handler, error) {
	fullSchemaContent, err := buildSchemaJSON(ctx, schemaFS, schemaFilePath, endpoint, true)
	if err != nil {
		return nil, fmt.Errorf("prepare full openapi schema JSON: %w", err)
	}

	schemaContent, err := buildSchemaJSON(ctx, schemaFS, schemaFilePath, endpoint, false)
	if err != nil {
		return nil, fmt.Errorf("prepare openapi schema JSON: %w", err)
	}

	handler := mux.NewRouter()

	// ручка для выдачи openapi полной схемы
	handler.HandleFunc(fullSchemaFileRoute, jsonContentHandleFunc(fullSchemaContent))

	// ручка для выдачи openapi схемы по файлам
	schemaFileRoute := fmt.Sprintf("%s/%s", schemaDirRoute, schemaFilePath)
	fileServerHandler := http.StripPrefix(schemaDirRoute, http.FileServer(http.FS(schemaFS)))
	schemaHandleFunc := jsonContentHandleFunc(schemaContent)
	handler.PathPrefix(schemaDirRoute).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL != nil && r.URL.Path == schemaFileRoute {
			schemaHandleFunc(w, r)
			return
		}

		fileServerHandler.ServeHTTP(w, r)
	})

	swaggerHandler := httpSwagger.Handler(httpSwagger.URL(schemaFileRoute))
	handler.PathPrefix("/").Handler(swaggerHandler)

	return handler, nil
}

func buildSchemaJSON(ctx context.Context, openapiSpecFS embed.FS, schemaFilePath string, endpoint string, withInternalizeRefs bool) ([]byte, error) {
	openapiSchemaContent, err := openapiSpecFS.ReadFile(schemaFilePath)
	if err != nil {
		return nil, fmt.Errorf("load openapi schema file: %w", err)
	}

	openapiLoader := openapi3.NewLoader()
	openapiLoader.IsExternalRefsAllowed = true
	openapiLoader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		return openapiSpecFS.ReadFile(url.Path)
	}

	openapi3Doc, err := openapiLoader.LoadFromData(openapiSchemaContent)
	if err != nil {
		return nil, fmt.Errorf("load openapi schema: %w", err)
	}

	server := &openapi3.Server{
		URL: endpoint,
	}
	openapi3Doc.AddServer(server)

	if withInternalizeRefs {
		// смерживаем все ссылки в одну схему
		openapi3Doc.InternalizeRefs(ctx, nil)
	}

	fullSchemaContent, err := openapi3Doc.MarshalJSON()
	if err != nil {
		return nil, fmt.Errorf("marshal schema to JSON: %w", err)
	}

	return fullSchemaContent, nil
}

func jsonContentHandleFunc(content []byte) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)

		_, writeErr := w.Write(content)
		if writeErr != nil {
			return
		}
	}
}
