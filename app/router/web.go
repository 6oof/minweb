package router

import (
	"github.com/6oof/minweb/app/middleware"
	"github.com/6oof/minweb/views/pages"
	"github.com/go-chi/chi/v5"
)

func RegisterWebRoutes(r *chi.Mux) {
	r.Group(func(r chi.Router) {
		r.Use(middleware.Csrf())

		// register all web routes below
		r.Get("/", pages.HandleIndex)

		//404 route
		r.Get("/*", pages.HandleNotFound)
	})

}
