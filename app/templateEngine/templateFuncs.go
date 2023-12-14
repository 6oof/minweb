package minitemp

import (
	"html/template"

	templatefuncs "github.com/6oof/chewbie/app/templateEngine/funcs"
)

var CustomTemplateFuncs = template.FuncMap{
	"customUppercase": templatefuncs.CustomUppercase,
}
