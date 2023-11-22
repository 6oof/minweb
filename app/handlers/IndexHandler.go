package handlers

import (
	"net/http"

	"github.com/6oof/chewbie/app/helpers"
	_ "github.com/mattn/go-sqlite3"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	// render the template
	html := helpers.RenderPage([]string{"layout", "test"}, "base", map[string]interface{}{
		"Name": "Supertiny",
	})

	// Send the HTML response
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(html))
}
