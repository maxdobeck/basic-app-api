package gatekeeper

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/sessions"
	"net/http"
)

type memberDetails struct {
	ID string
}

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

// ValidSession checks if the session is authenticated and still active
func ValidSession(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "scheduler-session")
	if err != nil {
		panic(err)
	}
	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Error(w, "Is this session valid: false", http.StatusUnauthorized)
		return
	}
	fmt.Fprintln(w, "Is this session valid: true")
}

// Login gets a new session for the user if the credential check passes
func Login(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "scheduler-session")
	if err != nil {
		fmt.Println(err)
	}
	creds := DecodeCredentials(r)
	// Authenticate based on incoming http request
	if passwordsMatch(r, creds) != true {
		http.Error(w, "Incorrect username or password", http.StatusUnauthorized)
		return
	}
	// Get the memberID based on the supplied email
	memberID := getMemberID(creds.Email)
	m := memberDetails {
		ID: memberID,
	}

	// Respond with the proper content type and the memberID
	w.Header().Set("Content-Type", "application/json") // TODO convert this to application/json
	// Set cookie values and save
	session.Values["authenticated"] = true
	session.Save(r, w)
	json.NewEncoder(w).Encode(m)
	// w.Write([]byte(memberID)) // Alternative to fprintf
}

// Logout destroys the session
func Logout(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "scheduler-session")
	if err != nil {
		panic(err)
	}
	// Revoke users authentication
	session.Values["authenticated"] = false
	session.Save(r, w)
}
