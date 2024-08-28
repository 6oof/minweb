package handlers

import (
	"net/http"

	"github.com/6oof/minweb/views/pages"
)

type NotFoundData struct {
	Code    int
	Message string
}

func HandleNotFound(w http.ResponseWriter, r *http.Request) {
	w.Write(pages.ErrorPage(404, "The page you are looking for does not exist.").Render())
}
