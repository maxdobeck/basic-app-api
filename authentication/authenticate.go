package authentication

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq" // github.com/lib/pq
	"net/http"
	"os"
)

// Credentials are the user provided email and password
type Credentials struct {
	Email, Password string
}

func getCurPassword(email string) (password string, userPresent bool) {
	connStr := os.Getenv("PGURL")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
	}
	sqlErr := db.QueryRow("SELECT password FROM members WHERE email = $1", email).Scan(&password)
	if sqlErr == sql.ErrNoRows {
		userPresent = false
		password = ""
		return
	}
	if sqlErr != nil {
		fmt.Println(sqlErr)
	}
	userPresent = true
	return
}

func passwordsMatch(r *http.Request, c Credentials) (match bool) {
	// c := DecodeCredentials(r)
	truePassword, userPresent := getCurPassword(c.Email)
	if userPresent != true {
		fmt.Println("User is not in the database")
		match = false
		return
	}
	if truePassword != c.Password {
		match = false
		fmt.Println("The passwords do not match")
		return
	}
	match = true
	return
}

// DecodeCredentials decodes the JSON data into a struct containing the email and password.DecodeCredentials
func DecodeCredentials(r *http.Request) (c Credentials) {
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		fmt.Println("Error decoding credentials >>", err)
	}
	return
}
