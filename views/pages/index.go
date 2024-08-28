package pages

import (
	"net/http"

	"github.com/6oof/minweb/app/helpers"
	"github.com/6oof/minweb/views"
	"github.com/6oof/minweb/views/fragments"
	"github.com/6oof/minweb/views/layouts"
	"github.com/6oof/xxhtml/x"
)

func IndexPage(r *http.Request, fdata fragments.ShowcaseFormResult) x.Elem {
	return layouts.Layout(views.SuperGlopabls(helpers.BaseSeo()),
		x.Section(`class="bg-white dark:bg-gray-900 text-center mt-16"`,
			x.Div(`class="grid max-w-screen-xl px-4 py-8 mx-auto"`,
				x.Div(`class="m-auto place-self-center "`,
					x.H1(`class="max-w-2xl mb-4 text-4xl font-extrabold tracking-tight leading-none md:text-5xl xl:text-6xl dark:text-primary"`,
						x.C(`MinWeb`),
					), x.P(`class="max-w-2xl mb-6 font-light text-gray-500 lg:mb-8 md:text-lg lg:text-xl dark:text-gray-400"`,
						x.C(`Toolkit for building oldschool websites.`),
					), x.A(`href="#" class="inline-flex items-center justify-center px-5 py-3 mr-3 text-base font-medium text-center text-white rounded-lg bg-primary-700 hover:bg-primary-800 focus:ring-4 focus:ring-primary-300 dark:focus:ring-primary-900"`,
						x.C(`Find Documentation`),
						x.E("svg", `class="w-5 h-5 ml-2 -mr-1" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"`,
							x.E("path", `fill-rule="evenodd" d="M10.293 3.293a1 1 0 011.414 0l6 6a1 1 0 010 1.414l-6 6a1 1 0 01-1.414-1.414L14.586 11H3a1 1 0 110-2h11.586l-4.293-4.293a1 1 0 010-1.414z" clip-rule="evenodd"`),
						),
					),
					fragments.ShowcaseForm(r, fdata),
				),
			),
		),
	)
}

// HandleIndex is the handler function for the "/" route, rendering the home page.
func HandleIndex(w http.ResponseWriter, r *http.Request) {

	fdata := fragments.ShowcaseFormResult{
		NameError: "",
		Result:    "",
	}

	w.Write(IndexPage(r, fdata).Render())
}
