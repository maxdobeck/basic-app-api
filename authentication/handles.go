//Package authentication challenges user credentials and creates or destroys cookie based sessions.
package authentication

import (
	"encoding/json"
	"github.com/antonlindstrom/pgstore"
	"github.com/gorilla/csrf"
	"github.com/maxdobeck/basic-app-api/models"
	"github.com/maxdobeck/basic-app-api/sessions"
	"log"
	"net/http"
	"os"
)

// Remove Member Details and errorMessage to other packages
type memberDetails struct {
	Status string
	Name   string
	Email  string
	ID     string
}

type errorMessage struct {
	Status  string
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

// Login gets a new session for the user if the credential check passes
func Login(w http.ResponseWriter, r *http.Request) {
	store, err := pgstore.NewPGStore(os.Getenv("PGURL"), key)
	check(err)
	defer store.Close()
	session, err := store.Get(r, "scheduler-session")
	check(err)

	var domain string
	if os.Getenv("GO_ENV") == "dev" {
		domain = "127.0.0.1:3000"
	} else if os.Getenv("GO_ENV") == "test" {
		domain = "http://s3-sih-test.s3-website-us-west-1.amazonaws.com"
	} else if os.Getenv("GO_ENV") == "prod" {
		domain = "https://schedulingishard.com"
	}
	// Limit the sessions to 1 24-hour day
	session.Options.MaxAge = 86400 * 1
	session.Options.Domain = domain // Set to localhost for testing only.  prod must be set to "schedulingishard.com"
	session.Options.HttpOnly = true

	creds := DecodeCredentials(r)
	// Authenticate based on incoming http request
	if passwordsMatch(r, creds) != true {
		log.Printf("Bad password for member: %v", creds.Email)
		msg := errorMessage{
			Status:  "Failed to authenticate",
			Message: "Incorrect username or password",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		//http.Error(w, "Incorrect username or password", http.StatusUnauthorized)
		json.NewEncoder(w).Encode(msg)
		return
	}
	// Get the memberID based on the supplied email
	memberID := models.GetMemberID(creds.Email)
	memberName := models.GetMemberName(memberID)
	m := memberDetails{
		Status: "OK",
		ID:     memberID,
		Name:   memberName,
		Email:  creds.Email,
	}

	// Respond with the proper content type and the memberID
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-CSRF-Token", csrf.Token(r))
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
	if sessions.GoodSession(r) != true {
		json.NewEncoder(w).Encode("Session Expired.  Log out and log back in.")
	}
	store, err := pgstore.NewPGStore(os.Getenv("PGURL"), key)
	check(err)
	defer store.Close()

	session, err := store.Get(r, "scheduler-session")
	check(err)
	// Revoke users authentication
	session.Values["authenticated"] = false
	session.Options.MaxAge = -1
	session.Save(r, w)
}

// CsrfToken will generate a CSRF Token
func CsrfToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-CSRF-Token", csrf.Token(r))
}
