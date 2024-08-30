package router

import (
	"net/http"

	"github.com/6oof/minweb/app/middleware"
	"github.com/go-chi/chi/v5"
)

func Base() *chi.Mux {
	r := chi.NewRouter()

	// Middleware setup
	r.Use(middleware.Logger())
	r.Use(middleware.Recoverer())
	r.Use(middleware.Cors())

	// Static file serving
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Register custom routes
	RegisterWebRoutes(r)
	RegisterApiRoutes(r)
	RegisterFragmentRoutes(r)

	return r
}
