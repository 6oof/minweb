package main

import (
	"github.com/6oof/chewbie/app/stubs"
	"github.com/6oof/chewbie/db"
)

func main() {
	db.InitDB("hello.db")
	stubs.MbinServe(":3334")
}
