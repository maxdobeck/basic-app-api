package gatekeeper

import (
	"fmt"
	"net/http"
	"github.com/gorilla/sessions"
)

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

// ValidSession checks if the session is authenticated and still active
func ValidSession(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "scheduler-session")
	
	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Error(w, "Is this session valid: false", http.StatusUnauthorized)
		return
	}

	// Return message
	fmt.Fprintln(w, "Is this session valid: true")
}

// Login gets a new session for the user if the credential check passes
func Login(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "scheduler-session")

	// Authenticate based on incoming http request
	/*
	*
	*
	*/

	// Set user as authenticated
	session.Values["authenticated"] = true
	session.Save(r, w)
}

// Logout destroys the session
func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "scheduler-session")

	// Revoke users authentication
	session.Values["authenticated"] = false
	session.Save(r, w)
}
