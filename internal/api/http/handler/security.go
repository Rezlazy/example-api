package handler

import (
	"context"
	"crypto/rsa"
	"example-api/internal/api/http/server"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	jwt.RegisteredClaims
	Email string
}

type SecurityHandler struct {
	publicKey *rsa.PublicKey
}

func NewSecurityHandler(publicKey *rsa.PublicKey) SecurityHandler {
	return SecurityHandler{
		publicKey: publicKey,
	}
}

func (sh SecurityHandler) HandleBearerAuth(ctx context.Context, operationName string, t server.BearerAuth) (context.Context, error) {
	claims := Claims{}
	_, err := jwt.ParseWithClaims(t.Token, &claims, func(token *jwt.Token) (interface{}, error) {
		return sh.publicKey, nil
	})
	if err != nil {
		return nil, fmt.Errorf("parse claims from token: %w", err)
	}

	return ctx, nil
}
