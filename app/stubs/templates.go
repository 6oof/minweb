package stubs

import (
	"bytes"
	"html/template"
	"strings"
)

// Define a custom template function that you want to register.
// example: This function takes a string and returns its uppercase version.
func customUppercase(str string) string {
	return strings.ToUpper(str)
}

var customTemplateFuncs = template.FuncMap{
	"customUppercase": customUppercase,
}

func RenderPage(files []string, name string, data interface{}) string {
	for i, v := range files {
		files[i] = "./templates/" + v + ".go.html"
	}
	sg := "./templates/superglobals.go.html"
	completeFiles := append([]string{sg}, files...)

	// Create a new template instance
	ts := template.New("")

	// Register your custom template function
	ts.Funcs(customTemplateFuncs)

	// Parse the template files
	ts, err := ts.ParseFiles(completeFiles...)
	if err != nil {
		panic(err)
	}

	var contentBuffer bytes.Buffer
	err = ts.ExecuteTemplate(&contentBuffer, name, data)

	return contentBuffer.String()
}

func RenderFragment(files []string, name string, data interface{}) string {
	for i, v := range files {
		files[i] = "./templates/" + v + ".go.html"
	}

	// Create a new template instance
	ts := template.New("")

	// Register your custom template function
	ts.Funcs(customTemplateFuncs)

	// Parse the template files
	ts, err := ts.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	var contentBuffer bytes.Buffer
	err = ts.ExecuteTemplate(&contentBuffer, name, data)

	return contentBuffer.String()
}
