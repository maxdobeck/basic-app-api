package sessions

import (
	"github.com/gorilla/csrf"
	"log"
	"net/http"
)

// CsrfToken will generate a CSRF Token
func CsrfToken(w http.ResponseWriter, r *http.Request) {
	log.Println("Generating csrf token")
	w.Header().Set("X-CSRF-Token", csrf.Token(r))
}

// ValidSession checks that the session is valid and can user can make requests
func ValidSession(w http.ResponseWriter, r *http.Request) {
	if GoodSession(r) != true {
		log.Println("Session is old, must log out log back in.")
		//w.WriteHeader(http.StatusUnauthorized)
		http.Error(w, "Session is expired.", http.StatusUnauthorized)
	} else {
		log.Println("Session is good.")
		w.WriteHeader(http.StatusOK)
	}
}
