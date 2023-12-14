package app

import (
	"github.com/6oof/chewbie/app/handlers"
	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r *chi.Mux) {
	r.Get("/", handlers.HandleIndex)
}
