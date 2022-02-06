package sessions

import (
	"github.com/gorilla/sessions"
	"log"
)

var SESSION_STORE *sessions.CookieStore

// GetSessionStoreFactory using Cookie as Session Store
func GetSessionStoreFactory() *sessions.CookieStore {
	if SESSION_STORE != nil {
		return SESSION_STORE
	}
	SESSION_STORE = sessions.NewCookieStore([]byte("t0p-s3cr3t"))
	log.Println("Initializing new Session Store!.")
	return SESSION_STORE
}
