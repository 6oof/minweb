package app

import (
	"github.com/6oof/chewbie/app/handlers"
	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/", handlers.HandleIndex).Methods("GET")

}
