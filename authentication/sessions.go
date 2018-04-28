package gatekeeper

import (
	"encoding/json"
	"github.com/antonlindstrom/pgstore"
	"log"
	"net/http"
	"os"
)

type memberDetails struct {
	Status string
	ID string
}

type errorMessage struct {
	Status string
	Message string
}

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

// ValidSession checks if the session is authenticated and still active
func ValidSession(w http.ResponseWriter, r *http.Request) {
	store, err := pgstore.NewPGStore(os.Getenv("PGURL"), []byte("secret-key"))
	check(err)
	defer store.Close()

	session, err := store.Get(r, "scheduler-session")
	check(err)

	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Error(w, "Invalid Session", http.StatusUnauthorized)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(auth)
		return
	}

	// Respond with the proper content type and the memberID
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("true")
	log.Println(w, "Is this session valid: true")
}

// Login gets a new session for the user if the credential check passes
func Login(w http.ResponseWriter, r *http.Request) {
	store, err := pgstore.NewPGStore(os.Getenv("PGURL"), []byte("secret-key"))
	check(err)
	defer store.Close()
	session, err := store.Get(r, "scheduler-session")
	check(err)
	// Limit the sessions to 1 24-hour day
	session.Options.MaxAge = 86400 * 1

	creds := DecodeCredentials(r)
	// Authenticate based on incoming http request
	if passwordsMatch(r, creds) != true {
		msg := errorMessage{
			Status: "Failed to authenticate",
			Message: "Incorrect username or password",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		//http.Error(w, "Incorrect username or password", http.StatusUnauthorized)
		json.NewEncoder(w).Encode(msg)
		return
	}
	// Get the memberID based on the supplied email
	memberID := getMemberID(creds.Email)
	m := memberDetails{
		Status: "OK",
		ID: memberID,
	}

	// Respond with the proper content type and the memberID
	w.Header().Set("Content-Type", "application/json")
	// Set cookie values and save
	session.Values["authenticated"] = true
	if err = session.Save(r, w); err != nil {
		log.Printf("Error saving session: %v", err)
	}
	json.NewEncoder(w).Encode(m)
	// w.Write([]byte(memberID)) // Alternative to fprintf
}

// Logout destroys the session
func Logout(w http.ResponseWriter, r *http.Request) {
	store, err := pgstore.NewPGStore(os.Getenv("PGURL"), []byte("secret-key"))
	check(err)
	defer store.Close()

	session, err := store.Get(r, "scheduler-session")
	check(err)
	// Revoke users authentication
	session.Values["authenticated"] = false
	session.Options.MaxAge = -1
	session.Save(r, w)
}
