package services

import (
	"github.com/gorilla/sessions"
)

// Store is an interface for custom session stores.
//
// See CookieStore and FilesystemStore for examples.

func MakeCookieSessionStore(key string) *sessions.CookieStore {
	return sessions.NewCookieStore([]byte(key))
}
