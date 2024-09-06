package app

import (
	"github.com/6oof/minweb/app/kernel"
	"github.com/6oof/minweb/database"
	"github.com/go-chi/chi/v5"
	"github.com/uptrace/bun"
)

// >>>>>>>>>>>>>>>>>>>>
//    /\  /\\  /\  /
//   /  \/  \\/  \/
// <<<<<<<<<<<<<<<<<<<<
// Service Locator
//
// general Info:
//
// 1. Fields in the `application` struct that use kernel interfaces (e.g., logger, storage) should generally not be switched out
//    unless absolutely necessary. This ensures consistency and reliability of core components.
// 2. To add additional services to the framework:
//    a. Add a new field to the `application` struct for the service.
//    b. Decide whether the service should be initialized in the `Boot` method or the `Start` method, depending on its nature.
//    c. Add a corresponding facade method to provide access to this new service in a controlled manner.
// 3. If multiple instances of a service are required (e.g., different loggers), consider creating separate services that adhere
//    to the same interface. For example, `storage` has a `privateStorage` counterpart to handle different types of storage.
//
// All services can be found in app/services and should never depend on the app package.

var app *application

// application container
type application struct {
	mux            *chi.Mux
	database       *bun.DB
	configs        kernel.ConfigInterface
	logger         kernel.LoggerInterface
	sessionStore   kernel.StoreInterface
	storage        kernel.StorageInterface
	privateStorage kernel.StorageInterface
	booted         bool // Indicates if the application has been initialized
}

// Boot initializes the core components of the application.
//
// General Instructions:
// 1. The order of initialization is crucial. Initialize configuration settings first, as they are required by other components.
// 2. Services that depend on the configuration should be initialized after the configuration is set up (e.g., logger).
// 3. Initialize components in the sequence they depend on one another to ensure proper setup (e.g., logger -> database -> storage).
// Ensure to follow this structure to maintain consistency and avoid issues during initialization.

func Boot() {
	app = &application{}

	cf := &kernel.Config{}
	app.configs = cf.InitConfig()

	appLogger := &kernel.AppLogger{}
	appLogger.Boot(app.configs.GetOrPanic("LOGGER_FILE"))
	app.logger = appLogger

	app.database = database.GetDb()

	app.sessionStore = kernel.MakeCookieSessionStore(app.configs.Get("key"))

	storage := &kernel.LocalStorage{}
	storage.Init(app.configs.GetOrPanic("STORAGE_PATH"))
	app.storage = storage

	privateStorage := &kernel.LocalStorage{}
	privateStorage.Init(app.configs.GetOrPanic("PRIVATE_STORAGE_PATH"))
	app.privateStorage = privateStorage

	app.booted = true
}

// Start begins the application's HTTP server
func Start(r *chi.Mux) {
	if !app.booted {
		panic("App not booted. Ensure that Boot() is called before Start().")
	}

	app.mux = r

	kernel.Serve(app.mux, ":"+app.configs.GetOrPanic("PORT"))
}

// Facade methods for accessing core components

// Log provides access to the logger for logging messages from anywhere in the app
func Log() kernel.LoggerInterface {
	return app.logger
}

// Config provides access to configuration values
func Config() kernel.ConfigInterface {
	return app.configs
}

// Database provides access to the database connection instance
func Database() *bun.DB {
	return app.database
}

// SessionStore provides access to the session store for managing user sessions
func SessionStore() kernel.StoreInterface {
	return app.sessionStore
}

// Storage provides access to the general storage instance for file storage
func Storage() kernel.StorageInterface {
	return app.storage
}

// PrivateStorage provides access to the private storage instance for sensitive files
func PrivateStorage() kernel.StorageInterface {
	return app.privateStorage
}
