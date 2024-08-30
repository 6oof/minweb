package main

import (
	"github.com/6oof/minweb/app"
	"github.com/6oof/minweb/router"
)

// Uncomment line below to use the default database file (1/2)
// db "github.com/6oof/minweb/database"

func main() {
	// Initialize the database with the specified file
	// uncomment line line below to use the default database file (2/2)
	// db.InitDB("mw.db")
	// Follow instructions in the documentation to migrate and generate queries

	// Boot the app
	app.Boot()

	// Start the app instance
	app.Get().Start(router.Base())
}
