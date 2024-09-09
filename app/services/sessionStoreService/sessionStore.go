package sessionStoreService

import (
	"github.com/gorilla/sessions"
)

// Store is an interface for custom session stores.
//
// See CookieStore and FilesystemStore for examples.

func Make(key string) *sessions.CookieStore {
	return sessions.NewCookieStore([]byte(key))
}
