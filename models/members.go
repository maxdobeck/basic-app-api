package models

import (
	_ "github.com/lib/pq" // github.com/lib/pq
	"log"
)

// NewMember is the struct for the member signup process
type NewMember struct {
	Name, Email, Email2, Password, Password2 string
}

// CreateMember creates the new member record
func CreateMember(m *NewMember) error {
	_, err := db.Query("INSERT INTO members(name, email, password) VALUES ($1,$2, $3)", m.Name, m.Email, m.Password)
	if err != nil {
		log.Println(err)
		return err
	}
	return err
}
