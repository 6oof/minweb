package middleware

import (
	"fmt"
	"net/http"

	"github.com/6oof/minweb/app"
	"github.com/gorilla/csrf"
)

func Csrf() func(next http.Handler) http.Handler {
	fmt.Println(app.Get().Config().GetOrPanic("KEY"))
	// CSRF protection
	key := []byte(app.Get().Config().GetOrPanic("KEY"))
	if len(key) < 1 {
		panic("App key must be set in .env file")
	}
	return csrf.Protect([]byte(key),
		csrf.Secure(false),
		csrf.ErrorHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte(`{"message": "Forbidden - CSRF token invalid"}`))
		}))) // Set Secure to true in production}
}
