package layouts

import "github.com/6oof/xxhtml/x"

func Empty(SuperGlobals x.Elem, Content x.Elem) x.Elem {
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
				),
			),
		),
	)
}
