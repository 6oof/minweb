package middleware

import (
	"net/http"

	"github.com/6oof/minweb/app/kernel/kernelmiddleware"
)

// Returns a nice error page with information about the panic and a stack trace.
func DevPanicPrint(next http.Handler) http.Handler {
	return kernelmiddleware.HtmlPanicRecovery(next)
}
