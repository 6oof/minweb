package views

import (
	"github.com/6oof/minweb/app/helpers"
	"github.com/6oof/xxhtml/x"
)

func SuperGlopabls(seo helpers.Seo) x.Elem {
	return x.E(
		"",
		x.Title(
			x.C(seo.Name+x.SIF(seo.Title != "", " • "+seo.Title)),
		),
		x.Meta(x.Att("name", "description"), x.Att("content", x.SIF(seo.Description != "", seo.Description))),
		x.IF(seo.Keywords != "",
			x.Meta(
				x.Att("name", "keywords"),
				x.Att("content", seo.Keywords),
			),
		),
		x.Link(
			x.Att("rel", "icon"),
			x.Att("type", "image/x-icon"),
			x.Att("href", "/static/assets/favicon.png"),
		),
		x.Link(
			x.Att("rel", "stylesheet"),
			x.Att("href", "/static/styles/main.css"),
		),
		x.Script(
			x.Att("defer", ""),
			x.Att("type", "module"),
			x.Att("src", "/static/scripts/main.js"),
		),
	)

}
