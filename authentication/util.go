package gatekeeper

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// Credentials are the user provided email and password
type Credentials struct {
	Email, Password string
}

func getMemberID(email string) (memberID string) {
	connStr := os.Getenv("PGURL")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
	}
	sqlErr := db.QueryRow("SELECT id FROM members WHERE email = $1", email).Scan(&memberID)
	if sqlErr == sql.ErrNoRows {
		memberID = ""
		return
	}
	if sqlErr != nil {
		fmt.Println(sqlErr)
	}
	return
}

// DecodeCredentials decodes the JSON data into a struct containing the email and password
func DecodeCredentials(r *http.Request) (c Credentials) {
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		fmt.Println("Error decoding credentials >>", err)
	}
	return
}
