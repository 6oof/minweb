package router

import (
	"github.com/6oof/minweb/app/handlers"
	"github.com/go-chi/chi/v5"
)

// registerRoutes sets up the routing for the MiniWeb application by defining the routes and associating them with their respective handlers.
func RegisterApiRoutes(r *chi.Mux) {
	r.Group(func(r chi.Router) {
		r.Route("/api", func(r chi.Router) {

			//register all api routes below
			r.Get("/", handlers.HandleIndex)

			// 404 route
			r.Get("/*", handlers.HandleNotFound)
		})
	})

}
