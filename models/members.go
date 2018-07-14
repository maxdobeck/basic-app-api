package models

import (
	"database/sql"
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
	_, err := Db.Query("INSERT INTO members(name, email, password) VALUES ($1,$2, $3)", m.Name, m.Email, hashedPw)
	if err != nil {
		log.Println(err)
		return err
	}
	return err
}

// GetMemberID uses the primary email of a user to get the memberID from the member's table
func GetMemberID(email string) (memberID string) {
	sqlErr := Db.QueryRow("SELECT id FROM members WHERE email = $1", email).Scan(&memberID)
	if sqlErr == sql.ErrNoRows {
		memberID = ""
		return
	}
	if sqlErr != nil {
		log.Println(sqlErr)
	}
	return
}

// GetMemberName grabs the name using the email
func GetMemberName(id string) (name string) {
	sqlErr := Db.QueryRow("SELECT name FROM members WHERE id =$1", id).Scan(&name)
	if sqlErr == sql.ErrNoRows {
		name = ""
		return
	}
	if sqlErr != nil {
		log.Println(sqlErr)
	}
	return name
}

// UpdateMemberName uses the member ID to insert a new name
func UpdateMemberName(id string, name string) bool {
	_, sqlErr := Db.Query("UPDATE members SET name = $2 WHERE id = $1", id, name)
	if sqlErr == sql.ErrNoRows {
		name = ""
		return false
	}
	if sqlErr != nil {
		log.Println(sqlErr)
	}
	return true
}

// UpdateMemberEmail uses the member ID to insert a new email
func UpdateMemberEmail(id string, email string) bool {
	_, sqlErr := Db.Query("UPDATE members SET name = $2 WHERE id = $1", id, email)
	if sqlErr == sql.ErrNoRows {
		return false
	}
	if sqlErr != nil {
		log.Println(sqlErr)
	}
	return true
}
