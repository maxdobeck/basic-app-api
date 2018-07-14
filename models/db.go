// Package models holds all the data layer interfaces and transactions
package models

import (
	"database/sql"
	_ "github.com/lib/pq" // github.com/lib/pq
	"log"
)

// Db is global variable pointer to database connection
var Db *sql.DB

// ConnToDB connects the database.  This should be called at the app start
func ConnToDB(dbURL string) {
	var err error
	Db, err = sql.Open("postgres", dbURL)
	if err != nil {
		log.Println(err)
	}
	if err = Db.Ping(); err != nil {
		log.Println(err)
	}
}
