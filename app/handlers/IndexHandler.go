package handlers

import (
	"net/http"

	minitemp "github.com/6oof/chewbie/app/templateEngine"
	_ "github.com/mattn/go-sqlite3"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	t := minitemp.PageTemplate{
		Files: []string{"layout", "index/index"},
		Data:  nil,
	}

	t.RenderPageAndSend(w)
}
