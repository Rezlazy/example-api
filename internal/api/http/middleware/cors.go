package middleware

import (
	"github.com/gorilla/handlers"
	"net/http"
)

const (
	corsOriginMatchAll = "*"
)

func CORSMW(h http.Handler) http.Handler {
	originsOk := handlers.AllowedOrigins([]string{
		corsOriginMatchAll,
	})

	headersOk := handlers.AllowedHeaders([]string{
		"Accept",
		"Accept-Language",
		"Content-Language",
		"Origin",
		"Content-Type",
		"Dnt",
		"Referer",
		"Authorization",
	})

	methodsOk := handlers.AllowedMethods([]string{
		http.MethodGet, http.MethodHead, http.MethodPost, http.MethodPatch, http.MethodOptions, http.MethodDelete,
	})

	return handlers.CORS(originsOk, headersOk, methodsOk)(h)
}
