package helpers

import (
	"github.com/6oof/minweb/app"
	"github.com/6oof/minweb/views"
)

// BaseSeo returns a default Seo object with values retrieved from environmental variables.

func BaseSeo() views.Seo {
	return views.Seo{
		Name:        app.Config().Get("NAME"),
		Description: app.Config().Get("DESCRIPTION"),
		// Title and Keywords can be added similarly using Env() for each.
		Ready: true,
	}
}
