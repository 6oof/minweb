package middleware

import (
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
)

func Compress() func(next http.Handler) http.Handler {
	return middleware.Compress(5, "text/html", "text/css")
}
