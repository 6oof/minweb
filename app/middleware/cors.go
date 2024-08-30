package middleware

import (
	"net/http"

	"github.com/6oof/minweb/app"
	"github.com/go-chi/cors"
)

func Cors() func(next http.Handler) http.Handler {

	// CORS protection
	return cors.Handler(cors.Options{
		AllowedOrigins:   []string{app.Config().Get("URL")}, // Use this to allow specific origin hosts
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
		// Debug:            true,
	})

}
