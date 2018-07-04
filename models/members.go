package models

import (
	_ "github.com/lib/pq" // github.com/lib/pq
	"golang.org/x/crypto/bcrypt"
	"log"
)

// NewMember is the struct for the member signup process
type NewMember struct {
	Name, Email, Email2, Password, Password2 string
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CreateMember creates the new member record
func CreateMember(m *NewMember) error {
	hashedPw, hashErr := hashPassword(m.Password)
	if hashErr != nil {
		log.Println("Error hashing password: ", hashErr)
	}
	_, err := db.Query("INSERT INTO members(name, email, password) VALUES ($1,$2, $3)", m.Name, m.Email, hashedPw)
	if err != nil {
		log.Println(err)
		return err
	}
	return err
}
