package router

import (
	"github.com/6oof/minweb/app/middleware"
	"github.com/6oof/minweb/views/fragments"
	"github.com/go-chi/chi/v5"
)

func RegisterFragmentRoutes(r *chi.Mux) {
	r.Group(func(r chi.Router) {
		r.Route("/!fragment", func(r chi.Router) {
			r.Use(middleware.Csrf())

			// Register all fragment routes below
			r.Post("/showcase-form", fragments.HandleShowcaseFormPost)

		})
	})

}
