package router

import (
	"net/http"

	"github.com/6oof/minweb/app"
	"github.com/6oof/minweb/app/middleware"
	"github.com/go-chi/chi/v5"
)

func ConstructRoutes() *chi.Mux {
	r := chi.NewRouter()

	// Middleware setup
	r.Use(middleware.Logger())
	r.Use(middleware.Recoverer())
	r.Use(middleware.Cors())
	r.Use(middleware.Compress())

	// Static file serving
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("storage/static"))))

	if app.Config().Get("STORAGE") == "local" {
		r.Handle("/static/*", http.StripPrefix("/publicstorage/", http.FileServer(http.Dir(app.Config().Get("STORAGE_PATH")))))
	}

	// Register custom routes
	RegisterWebRoutes(r)
	RegisterApiRoutes(r)
	RegisterFragmentRoutes(r)

	return r
}
