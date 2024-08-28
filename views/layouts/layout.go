package layouts

import (
	"time"

	"github.com/6oof/minweb/app/helpers"
	"github.com/6oof/xxhtml/x"
)

func Layout(SuperGlobals x.Elem, Content x.Elem) x.Elem {
	return x.E("", "",

		x.DOCTYPE(),

		x.Html(`lang="en"`,
			x.Head("",
				x.Meta(`charset="utf-8"`),
				x.Meta(`name="viewport" content="width=device-width, initial-scale=1"`),
				SuperGlobals,
			),
			x.Body(`class="bg-gray-900 min-h-screen flex flex-col justify-between"`,
				x.Main("",
					Content,
				), x.Footer(`class="bg-white rounded-lg shadow m-4 dark:bg-gray-800"`,
					x.Div(`class="w-full mx-auto max-w-screen-xl p-4 md:flex md:items-center md:justify-between"`,
						x.Span(`class="text-sm text-gray-500 sm:text-center dark:text-gray-400"`,
							x.C(time.Now().Year()),
							x.A(`href="/" class="hover:underline"`,
								x.C(helpers.Env("NAME", "MiniWeb")),
							),
							x.C(`. All Rights Reserved.`),
						),
						x.Ul(`class="flex flex-wrap items-center mt-3 text-sm font-medium text-gray-500 dark:text-gray-400 sm:mt-0"`,
							x.Li("",
								x.A(`href="#" class="hover:underline"`,
									x.C(`Github`),
								),
							),
						),
					),
				),
			),
		),
	)
}
