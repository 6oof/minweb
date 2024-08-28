package middleware

import (
	"net/http"

	cm "github.com/go-chi/chi/v5/middleware"
)

func Logger() func(next http.Handler) http.Handler {
	// CORS protection
	return cm.Logger
}
