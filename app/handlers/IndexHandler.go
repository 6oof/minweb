package handlers

import (
	"net/http"

	"github.com/6oof/minweb/views/components"
	"github.com/6oof/minweb/views/pages"
)

// HandleIndex is the handler function for the "/" route, rendering the home page.
func HandleIndex(w http.ResponseWriter, r *http.Request) {

	fdata := components.ShowcaseFormResult{
		NameError: "",
		Result:    "",
	}

	w.Write(pages.IndexPage(r, fdata).Render())
}

func HandleShowcaseFormPost(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()

	if err != nil {
		http.Error(w, "Sorry, something went wrong", http.StatusInternalServerError)
	}

	data := components.ShowcaseFormResult{}

	if len(r.PostForm.Get("Name")) < 1 {
		data.NameError = "Name is required"
	} else {
		if r.PostForm.Get("Name") == "Bruce Wayne" {
			data.Result = "You're Batman"
		} else {
			data.Result = "You're not Batman"
		}
	}

	w.Write(components.ShowcaseForm(r, data).Render())

}
