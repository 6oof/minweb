package app

import (
	"github.com/6oof/minweb/app/kernel"
	"github.com/go-chi/chi/v5"
	"github.com/uptrace/bun"
)

var app *application

func (app *application) Start(r *chi.Mux) {
	//set mux
	app.mux = r

	// Start the server
	kernel.Serve(app.mux, ":"+app.configs.GetOrPanic("PORT"))
}

func Get() *application {
	return app
}

// Log provides access to the logger for logging messages from anywhere in the app.
func Log() kernel.LoggerInterface {
	return app.logger
}

// Config provides access to configuration values.
func Config() *kernel.Config {
	return app.configs
}

func DB() *bun.DB {
	return app.database
}
