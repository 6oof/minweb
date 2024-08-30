package layouts

import (
	"time"

	"github.com/6oof/minweb/app"
	"github.com/6oof/xxhtml/x"
)

func Layout(SuperGlobals x.Elem, Content x.Elem) x.Elem {
	return x.E("",
		x.DOCTYPE(),
		x.Html(
			x.Att("lang", "en"),
			x.Head(
				x.Meta(x.Att("charset", "utf-8")),
				x.Meta(x.Att("name", "viewport"), x.Att("content", "width=device-width, initial-scale=1")),
				SuperGlobals,
			),
			x.Body(
				x.Class("bg-gray-900 min-h-screen flex flex-col justify-between"),
				x.Main(
					Content,
				),
				x.Footer(
					x.Class("bg-white rounded-lg shadow m-4 dark:bg-gray-800"),
					x.Div(
						x.Class("w-full mx-auto max-w-screen-xl p-4 md:flex md:items-center md:justify-between"),
						x.Span(
							x.Class("text-sm text-gray-500 sm:text-center dark:text-gray-400"),
							x.C(time.Now().Year()),
							x.A(
								x.Att("href", "/"),
								x.Class("hover:underline"),
								x.C(app.Config().Get("NAME")),
							),
							x.C(". All Rights Reserved."),
						),
						x.Ul(
							x.Class("flex flex-wrap items-center mt-3 text-sm font-medium text-gray-500 dark:text-gray-400 sm:mt-0"),
							x.Li(
								x.A(
									x.Att("href", "#"),
									x.Class("hover:underline"),
									x.C("Github"),
								),
							),
						),
					),
				),
			),
		),
	)

}
