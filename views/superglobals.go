package views

import (
	"github.com/6oof/xxhtml/x"
)

// Seo represents the SEO (Search Engine Optimization) metadata for a web page.
type Seo struct {
	// Name is the name of the web page.
	Name string

	// Title is the title of the web page.
	Title string

	// Description is the meta description for the web page.
	Description string

	// Keywords are the meta keywords for the web page.
	Keywords string

	// Ready is a flag indicating whether the SEO data is ready to be used.
	Ready bool
}

func SuperGlopabls(seo Seo) x.Elem {
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
