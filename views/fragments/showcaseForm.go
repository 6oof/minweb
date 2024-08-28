package fragments

import (
	"net/http"

	"github.com/6oof/minweb/views/components"
	"github.com/6oof/xxhtml/x"
)

type ShowcaseFormResult struct {
	NameError string
	Result    string
}

func ShowcaseForm(r *http.Request, fdata ShowcaseFormResult) x.Elem {
	return x.Div(`class="mt-8" id="result" hx-swap="outerHTML"`,
		x.Form(`class="max-w-sm mx-auto" hx-post="/!fragment/showcase-form" hx-target="#result"`,
			x.Div(`class="mb-5"`,
				components.CSRF(r),
				x.Label(`for="name" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white"`,
					x.C(`A quick and dirty example of HTMX Tailwind and MinWeb working together`),
				), x.Input(`type="text" id="name"`+x.STER(fdata.NameError != "",
					`class="bg-red-50 border border-red-500 text-red-900 placeholder-red-700 text-sm rounded-lg focus:ring-red-500 focus:border-red-500 block w-full p-2.5 dark:bg-red-100 dark:border-red-400"`,
					`class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary focus:border-primary block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary dark:focus:border-primary"`)+`placeholder="Bruce Wayne" name="Name"`),
				x.IF(fdata.NameError != "",
					x.Div(`class="text-red-700"`,
						x.C(fdata.NameError),
					),
				),
			),
			x.Button(`type="submit" class="text-white bg-primary hover:bg-primary focus:ring-4 focus:outline-none focus:ring-primary font-medium rounded-lg text-sm w-full sm:w-auto px-5 py-2.5 text-center dark:bg-primary dark:hover:bg-primary dark:focus:ring-primary"`,
				x.C(`Dispatch`),
			),
		),
		x.IF(fdata.Result != "",
			x.Div(`class="text-xl text-white mt-8"`,
				x.C(fdata.Result),
			),
		),
	)
}

func HandleShowcaseFormPost(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()

	if err != nil {
		http.Error(w, "Sorry, something went wrong", http.StatusInternalServerError)
	}

	data := ShowcaseFormResult{}

	if len(r.PostForm.Get("Name")) < 1 {
		data.NameError = "Name is required"
	} else {
		if r.PostForm.Get("Name") == "Bruce Wayne" {
			data.Result = "You're Batman"
		} else {
			data.Result = "You're not Batman"
		}
	}

	w.Write(ShowcaseForm(r, data).Render())

}
