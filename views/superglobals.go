package views

import (
	"github.com/6oof/minweb/app/helpers"
	"github.com/6oof/xxhtml/x"
)

func SuperGlopabls(seo helpers.Seo) x.Elem {
	return x.E("", "",
		x.Title("",
			x.C(seo.Name+x.SIF(seo.Title != "", ` • `+seo.Title)),
		),
		x.Meta(`name="description" content="`+x.SIF(seo.Description != "", seo.Description)+`"`),
		x.IF(seo.Keywords != "", x.Meta(`name="keywords" content="`+seo.Keywords+`"`)),
		x.Link(`rel="icon" type="image/x-icon" href="/static/assets/favicon.png"`),
		x.Link(`rel="stylesheet" href="/static/styles/main.css"`),
		x.Script(`defer="" type="module" src="/static/scripts/main.js"`),
	)
}
