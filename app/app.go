package app

import (
	"github.com/6oof/minweb/app/kernel"
	db "github.com/6oof/minweb/database"
	"github.com/go-chi/chi/v5"
	"github.com/uptrace/bun"
)

type application struct {
	mux      *chi.Mux
	logger   kernel.LoggerInterface
	configs  *kernel.Config
	database *bun.DB
}

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

	app.database = db.GetDb()
}
