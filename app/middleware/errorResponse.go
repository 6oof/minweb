package middleware

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime/debug"
	"strings"
	"text/template"
	"time"

	"github.com/6oof/minweb/app/helpers"
	"github.com/6oof/minweb/views"
	"github.com/6oof/minweb/views/layouts"
	"github.com/6oof/minweb/views/pages"
	"github.com/6oof/xxhtml/x"
)

//This is an absolute GPT fuled hackjob. Full credit goes to go-chi where all of this is copied from: https://github.com/go-chi/chi/blob/master/middleware/recoverer.go
// +
// The original work was derived from Goji's middleware, source:
// https://github.com/zenazn/goji/tree/master/web/middleware

// DevPanicPrint is a middleware that recovers from panics, and returns a detailed HTTP 500 (Internal Server Error) status
// with debug information in development, or a generic error page in production.
func DevPanicPrint(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rvr := recover(); rvr != nil {
				if rvr == http.ErrAbortHandler {
					// We don't recover http.ErrAbortHandler so the response
					// to the client is aborted; this should not be logged
					panic(rvr)
				}

				env := helpers.Env("ENVIROMENT", "prod")
				if env == "dev" {
					// Detailed error response for development environment
					w.WriteHeader(http.StatusInternalServerError)
					w.Header().Set("Content-Type", "text/html")
					stackTrace := string(debug.Stack())
					escapedStackTrace := HTMLStackTrace(rvr, []byte(stackTrace))
					errorPage := debugPage(escapedStackTrace, time.Now().Format(time.RFC3339), r.URL.String(), r.Method, rvr)
					w.Write(errorPage.Render())
				} else {
					// Generic error page for production environment
					w.WriteHeader(http.StatusInternalServerError)
					w.Header().Set("Content-Type", "text/html")
					ep := pages.ErrorPage(500, "Something went wrong on our end. Please try again later.")

					w.Write([]byte(ep.Render()))
				}
			}
		}()

		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}

// DebugPage generates a debug page element for rendering stack traces or other error information.
func debugPage(strace string, time string, url string, method string, rvr interface{}) x.Elem {
	return layouts.Empty(
		views.SuperGlopabls(helpers.BaseSeo()),
		x.E(
			"",
			x.Div(x.Class("divide-y divide-gray-200 overflow-hidden rounded-lg bg-white shadow p-6 m-8"),
				x.Div(x.Class("mb-4"),
					x.H1(x.Class("text-2xl font-semibold leading-6 text-red-900"),
						x.C("500 Internal Server Error"),
					),
				),
				x.Div(x.Class("pt-4"),
					x.Div(x.Class("details bg-gray-100 p-4 rounded-lg"),
						x.P(
							x.E(
								"strong",
								x.C("Panic:"),
							),
						),
						x.P(
							x.CR(template.HTMLEscapeString(fmt.Sprintf("%v", rvr))),
						),
						x.P(
							x.E(
								"strong",
								x.C("Request URL:"),
							),
							x.C(url),
						),
						x.P(
							x.E(
								"strong",
								x.C("Request Method:"),
							),
							x.C(method),
						),
						x.P(
							x.E(
								"strong",
								x.C("Timestamp:"),
							),
							x.C(time),
						),
						x.P(
							x.Class("mb-4"),
							x.E(
								"strong",
								x.C("Stack Trace:"),
							),
						),
						x.Div(x.Class("bg-gray-800 text-white rounded break-words overflow-x-auto"),
							x.ERAW(strace),
						),
					),
				),
			),
		),
	)
}

// HTMLStackTrace generates an HTML formatted stack trace with similar styling to the terminal output.
func HTMLStackTrace(rvr interface{}, debugStack []byte) string {
	s := htmlStack{}
	htmlOutput, err := s.parse(debugStack, rvr)
	if err != nil {
		// If there's an error, fall back to plain stack trace
		return `<pre>` + template.HTMLEscapeString(string(debugStack)) + `</pre>`
	}
	return htmlOutput
}

type htmlStack struct{}

func (s htmlStack) parse(debugStack []byte, rvr interface{}) (string, error) {
	var buf bytes.Buffer

	buf.WriteString(`<div style="font-family: Arial, sans-serif;  padding: 15px; border: 1px solid #ddd; border-radius: 4px;">`)
	buf.WriteString(`<pre>`)

	stack := strings.Split(string(debugStack), "\n")
	lines := []string{}

	// Locate panic line, as we may have nested panics
	for i := len(stack) - 1; i > 0; i-- {
		lines = append(lines, stack[i])
		if strings.HasPrefix(stack[i], "panic(") {
			lines = lines[0 : len(lines)-2] // Remove boilerplate
			break
		}
	}

	// Reverse lines
	for i := len(lines)/2 - 1; i >= 0; i-- {
		opp := len(lines) - 1 - i
		lines[i], lines[opp] = lines[opp], lines[i]
	}

	// Decorate lines
	for i, line := range lines {
		decoratedLine, err := s.decorateLine(line, i)
		if err != nil {
			return "", err
		}
		buf.WriteString(decoratedLine)
	}

	buf.WriteString(`</pre>`)
	buf.WriteString(`</div>`)

	return buf.String(), nil
}

func (s htmlStack) decorateLine(line string, num int) (string, error) {
	line = strings.TrimSpace(line)
	if strings.HasPrefix(line, "\t") || strings.Contains(line, ".go:") {
		return s.decorateSourceLine(line, num), nil
	}
	if strings.HasSuffix(line, ")") {
		return s.decorateFuncCallLine(line, num), nil
	}
	if strings.HasPrefix(line, "\t") {
		return `<div style="margin-left: 20px;">` + template.HTMLEscapeString(line) + `</div>`, nil
	}
	return `<div>` + template.HTMLEscapeString(line) + `</div>`, nil
}

func (s htmlStack) decorateFuncCallLine(line string, num int) string {
	idx := strings.LastIndex(line, "(")
	if idx < 0 {
		return ""
	}

	pkg := line[:idx]
	method := ""

	if idx := strings.LastIndex(pkg, "/"); idx < 0 {
		if idx := strings.Index(pkg, "."); idx > 0 {
			method = pkg[idx:]
			pkg = pkg[:idx]
		}
	} else {
		method = pkg[idx+1:]
		pkg = pkg[:idx+1]
		if idx := strings.Index(method, "."); idx > 0 {
			pkg += method[:idx]
			method = method[idx:]
		}
	}

	var colorStyle string
	if num == 0 {
		colorStyle = "color: #d9534f;"
	} else {
		colorStyle = "color: #5bc0de;"
	}

	return `<div style="margin-left: 20px; ` + colorStyle + `">` + template.HTMLEscapeString(pkg) + template.HTMLEscapeString(method) + `</div>`
}

func (s htmlStack) decorateSourceLine(line string, num int) string {
	idx := strings.LastIndex(line, ".go:")
	if idx < 0 {
		return ""
	}

	path := line[:idx+3]
	lineno := line[idx+3:]

	idx = strings.LastIndex(path, "/")
	dir := path[:idx+1]
	file := path[idx+1:]

	idx = strings.Index(lineno, " ")
	if idx > 0 {
		lineno = lineno[:idx]
	}
	return `<div style="margin-left: 20px; color: #f7f7f7;">` + template.HTMLEscapeString(dir) + template.HTMLEscapeString(file) + `:` + template.HTMLEscapeString(lineno) + `</div>`
}
