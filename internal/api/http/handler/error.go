package handler

import (
	"context"
	"example-api/internal/api/http/server"
	"example-api/pkg/logger"
	"net/http"
)

func (h Handler) NewError(ctx context.Context, err error) *server.ErrorResponseStatusCode {
	logger.Error(ctx, err)

	return &server.ErrorResponseStatusCode{
		StatusCode: http.StatusInternalServerError,
		Response: server.ErrorResponse{
			ErrorMessage: err.Error(),
		},
	}
}
