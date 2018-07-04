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
