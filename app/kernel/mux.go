package kernel

import (
	"github.com/go-chi/chi/v5"
)

// InitMux initializes the MiniWeb application, sets up middleware, and registers routes.
func InitMux() *chi.Mux {
	r := chi.NewRouter()

	return r
}
