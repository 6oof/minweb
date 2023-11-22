package main

import (
	"log"
	"net/http"

	"github.com/fridauxd/cht/app"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	fs := http.FileServer(http.Dir("static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))

	app.RegisterRoutes(r)

	port := ":3094"
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatalf("Error starting server: %v", err)
	} else {
		log.Printf("Server is running on port %s...", port)
	}

}
