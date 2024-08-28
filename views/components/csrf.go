package components

import (
	"net/http"

	"github.com/6oof/xxhtml/x"
	"github.com/gorilla/csrf"
)

func CSRF(r *http.Request) x.Elem {
	return x.ERAW(string(csrf.TemplateField(r)))
}
