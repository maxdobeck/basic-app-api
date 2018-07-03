package main

import (
	"github.com/antonlindstrom/pgstore"
	"github.com/gorilla/context"
	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
	"github.com/maxdobeck/gatekeeper/authentication"
	"github.com/maxdobeck/gatekeeper/members"
	"github.com/maxdobeck/gatekeeper/models"
	"github.com/maxdobeck/gatekeeper/sessions"
	"github.com/rs/cors"
	"github.com/urfave/negroni"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	store, err := pgstore.NewPGStore(os.Getenv("PGURL"), []byte("secret-key"))
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer store.Close()

	// Run a background goroutine to clean up expired sessions from the database.
	defer store.StopCleanup(store.Cleanup(time.Minute * 5))

	connStr := os.Getenv("PGURL")
	models.ConnToDB(connStr)

	var allowedDomains []string
	if os.Getenv("GO_ENV") == "dev" {
		allowedDomains = []string{"http://127.0.0.1:3030", "http://localhost:3030"}
	} else if os.Getenv("GO_ENV") == "test" {
		allowedDomains = []string{"http://s3-sih-test.s3-website-us-west-1.amazonaws.com"}
	} else if os.Getenv("GO_ENV") == "prod" {
		allowedDomains = []string{"https://schedulingishard.com", "https://www.schedulingishard.com"}
	}

	CSRF := csrf.Protect(
		[]byte("32-byte-long-auth-key"),
		csrf.RequestHeader("X-CSRF-Token"),
		csrf.CookieName("scheduler_csrf"),
		csrf.Secure(false), // Disabled for localhost non-https debugging
	)

	c := cors.New(cors.Options{
		AllowedOrigins:   allowedDomains,
		AllowCredentials: true,
		AllowedHeaders:   []string{"X-CSRF-Token"},
		ExposedHeaders:   []string{"X-CSRF-Token"},
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	})

	r := mux.NewRouter()
	// Authentication Routes
	r.HandleFunc("/csrftoken", sessions.CsrfToken).Methods("GET")
	r.HandleFunc("/login", authentication.Login).Methods("POST")
	r.HandleFunc("/logout", authentication.Logout).Methods("POST")
	// Member CRUD routes
	r.HandleFunc("/members", members.SignupMember).Methods("POST")
	// Middleware
	n := negroni.Classic()
	n.Use(c)
	n.UseHandler(CSRF(r))

	var hostURL string
	if os.Getenv("GO_ENV") == "test" {
		hostURL = "https://shielded-stream-75107.herokuapp.com/"
	} else if os.Getenv("GO_ENV") == "dev" {
		hostURL = "http://localhost"
	}
	port := os.Getenv("PORT")
	log.Println("Listening on: ", hostURL)
	log.Fatal(http.ListenAndServe(":"+port, context.ClearHandler(n)))
}
