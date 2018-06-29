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

	CSRF := csrf.Protect(
		[]byte("32-byte-long-auth-key"),
		csrf.RequestHeader("X-CSRF-Token"),
		csrf.CookieName("scheduler_csrf"),
		csrf.Secure(false), // Disabled for localhost non-https debugging
	)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://127.0.0.1:3030", "http://localhost:3030", "https://schedulingishard.com", "https://www.schedulingishard.com"},
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

	log.Println("Listening on http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", context.ClearHandler(n)))
}
