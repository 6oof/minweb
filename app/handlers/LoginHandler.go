package handlers

import (
	"net/http"

	"github.com/fridauxd/cht/app/helpers"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	html := helpers.RenderPage([]string{"layout", "auth/login", "auth/loginform"}, "base", map[string]interface{}{
		"Errors": nil,
	})

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(html))

}

type UserLogin struct {
	Login    string `form:"login" validate:"required,email"`
	Password string `form:"password" validate:"required"`
	Remember bool   `form:"remember"`
}

func HandleLoginPost(w http.ResponseWriter, r *http.Request) {
	var userLogin UserLogin

	validator := helpers.NewValidator()
	if err := validator.ValidateAndMapForm(w, r, &userLogin); err != nil {

		ee := validator.ExtractErrors(err)

		html := helpers.RenderFragment([]string{"auth/loginform"}, "formCore", map[string]interface{}{
			"Errors": ee,
		})

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(html))
	}

	w.Header().Set("HX-Redirect", "/")
	w.WriteHeader(http.StatusOK)
}
