package pages

import (
	"github.com/6oof/minweb/app/helpers"
	"github.com/6oof/minweb/views"
	"github.com/6oof/minweb/views/layouts"
	"github.com/6oof/xxhtml/x"
)

func ErrorPage(code int, message string) x.Elem {
	return layouts.Empty(views.SuperGlopabls(helpers.BaseSeo()),
		x.Section(`class="bg-white dark:bg-gray-900"`,
			x.Div(`class="py-8 px-4 mx-auto max-w-screen-xl lg:py-16 lg:px-6"`,
				x.Div(`class="mx-auto max-w-screen-sm text-center"`,
					x.H1(`class="mb-4 text-7xl tracking-tight font-extrabold lg:text-9xl text-primary dark:text-primary"`,
						x.C(code),
					),
					x.P(`class="mb-4 text-3xl tracking-tight font-bold text-gray-900 md:text-4xl dark:text-white"`,
						x.C(`Something's wrong.`),
					),
					x.P(`class="mb-4 text-lg font-light text-gray-500 dark:text-gray-400"`,
						x.C(message),
					),
					x.A(`href="/" class="inline-flex text-white bg-primary-600 hover:bg-primary-800 focus:ring-4 focus:outline-none focus:ring-primary-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:focus:ring-primary-900 my-4"`,
						x.C(`Back to Homepage`),
					),
				),
			),
		),
	)
}
