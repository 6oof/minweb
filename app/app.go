package app

import (
	"github.com/6oof/minweb/app/kernel"
	"github.com/go-chi/chi/v5"
)

type Application struct {
	mux     *chi.Mux
	logger  kernel.LoggerInterface
	configs *kernel.Config
}

var app *Application

// Boot initializes the application.
func Boot() {
	app = &Application{}
	//Initialize configuration
	config := kernel.InitConfig()
	app.configs = config

	// Create an instance of AppLogger
	appLogger := &kernel.AppLogger{}
	appLogger.Boot() // Boot the logger to initialize it
	app.logger = appLogger
}

func (app *Application) Start(r *chi.Mux) {
	app.mux = r

	// Start the server
	kernel.Serve(app.mux)
}

func Get() *Application {
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
