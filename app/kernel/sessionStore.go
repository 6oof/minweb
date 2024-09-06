package kernel

import (
	"github.com/gorilla/sessions"
)

// Store is an interface for custom session stores.
//
// See CookieStore and FilesystemStore for examples.
type StoreInterface sessions.Store

func MakeCookieSessionStore(key string) *sessions.CookieStore {
	return sessions.NewCookieStore([]byte(key))
}
