package mwtemp

import (
	"html/template"

	"github.com/6oof/miniweb-base/app/templateEngine/templatefuncs"
)

// CustomTemplateFuncs is a FuncMap that holds custom template functions used in rendering templates.
var CustomTemplateFuncs = template.FuncMap{
	"CurrentYear": templatefuncs.CurrentYear,
	"ProjectName": templatefuncs.ProjectName,
}
