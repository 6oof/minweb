package app

import (
	"github.com/6oof/minweb/app/kernel"
	"github.com/go-chi/chi/v5"
)

type application struct {
	mux     *chi.Mux
	logger  kernel.LoggerInterface
	configs *kernel.Config
}

var app *application

// Boot initializes the application.
func Boot() {
	app = &application{}
	//Initialize configuration
	config := kernel.InitConfig()
	app.configs = config

	// Create an instance of AppLogger
	appLogger := &kernel.AppLogger{}
	appLogger.Boot(config.GetOrPanic("LOGGER_FILE")) // Boot the logger to initialize it
	app.logger = appLogger
}

func (app *application) Start(r *chi.Mux) {
	//set mux
	app.mux = r

	// Start the server
	kernel.Serve(app.mux, app.configs.GetOrPanic("PORT"))
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
