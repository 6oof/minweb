package main

import (
	"github.com/6oof/minweb/app"
	"github.com/6oof/minweb/router"
)

func main() {
	app.Boot()
	app.Start(router.ConstructRoutes())
}
