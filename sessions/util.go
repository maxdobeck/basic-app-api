// Package sessions resolves session related issues
package sessions

import (
	"github.com/antonlindstrom/pgstore"
	_ "github.com/lib/pq" // github.com/lib/pq
	"log"
	"net/http"
	"os"
)

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key = []byte("super-secret-key")
	// store = sessions.NewCookieStore(key)
)

func check(err error) {
	if err != nil {
		log.Println(err)
	}
}

// GoodSession returns true or false depending on if the session is current
func GoodSession(r *http.Request) bool {
	store, err := pgstore.NewPGStore(os.Getenv("PGURL"), key)
	check(err)
	defer store.Close()

	session, err := store.Get(r, "scheduler-session")
	check(err)

	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		return false
	}
	return true
}
