package layouts

import "github.com/6oof/xxhtml/x"

func Empty(SuperGlobals x.Elem, Content x.Elem) x.Elem {
	return x.E(
		"",
		x.DOCTYPE(),
		x.Html(
			x.Att("lang", "en"),
			x.Head(
				x.Meta(x.Att("charset", "utf-8")),
				x.Meta(
					x.Att("name", "viewport"),
					x.Att("content", "width=device-width, initial-scale=1"),
				),
				SuperGlobals,
			),
			x.Body(
				x.Class("bg-gray-900 min-h-screen flex flex-col justify-between"),
				x.Main(
					Content,
				),
			),
		),
	)

}
