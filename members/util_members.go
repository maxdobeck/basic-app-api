package members

import (
	"database/sql"
	_ "github.com/lib/pq" // github.com/lib/pq
	"github.com/maxdobeck/basic-app-api/models"
	"log"
	"os"
)

func uniqueEmail(email string) bool {
	connStr := os.Getenv("PGURL")
	models.ConnToDB(connStr)
	log.Println("we're in it now boyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyssssssssssssssssssssssssssss")
	var existingEmail string
	row := models.Db.QueryRow("SELECT email FROM members WHERE email = $1", email)
	err := row.Scan(&existingEmail)
	if err != nil {
		log.Println(err)
	}
	if err == sql.ErrNoRows {
		log.Println("Email does not exist.")
		return true
	}

	log.Println("Email exists in the store: ", existingEmail)
	if len(existingEmail) > 0 {
		return false
	}
	return true
}

// emailsMatch returns true if both emails are the same
func emailsMatch(email1 string, email2 string) bool {
	if email1 == email2 {
		return true
	}
	return false
}

// emailAvailable returns true if the email is not taken
func emailAvailable(email string) bool {
	if uniqueEmail(email) == true {
		return true
	}
	return false
}

// passwordsMatch confirms that both passwords are matching.  This helps the user avoid typing the incorrect password
func passwordsMatch(pw1 string, pw2 string) bool {
	if pw1 == pw2 {
		return true
	}
	return false
}
