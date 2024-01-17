package handler

import (
	"context"
	"example-api/internal/api/http/server"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
)

const expiresAtIntervalDefault = 15

func (h Handler) V1AuthLoginPost(ctx context.Context, req *server.AuthLoginRequestSchema) (*server.Token, error) {
	if req.Email != h.email || req.Password != h.password {
		return nil, &server.ErrorResponseStatusCode{
			StatusCode: http.StatusUnauthorized,
		}
	}

	token := jwt.New(jwt.GetSigningMethod("RS256"))
	token.Claims = &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * expiresAtIntervalDefault)),
		},
		Email: req.Email,
	}

	accessToken, err := token.SignedString(h.privateKey)
	if err != nil {
		return nil, fmt.Errorf("signed token: %w", err)
	}

	return &server.Token{
		AccessToken: accessToken,
	}, nil
}
