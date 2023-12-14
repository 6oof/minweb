package main

import (
	"github.com/6oof/chewbie/app"
	db "github.com/6oof/chewbie/database"
)

func main() {
	db.InitDB("hello.db")
	app.MbinServe(":3033")
}
