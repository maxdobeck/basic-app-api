package models

import (
	_ "github.com/lib/pq" // github.com/lib/pq
	"log"
	"os"
	"testing"
)

func TestCreateMember(t *testing.T) {
	ConnToDB(os.Getenv("PGURL"))

	_, delErr := db.Query("DELETE FROM members WHERE email like 'testtest@gmail.com'")
	log.Println(delErr)

	m := NewMember{
		Name:      "Test Member",
		Email:     "testtest@gmail.com",
		Email2:    "testtest@gmail.com",
		Password:  "superduper",
		Password2: "superduper",
	}

	if CreateMember(&m) != nil {
		t.Fail()
	}

	var record string
	err := db.QueryRow("SELECT email FROM members WHERE email like 'testtest@gmail.com'").Scan(&record)
	if err != nil {
		log.Println(err)
		t.Log(err)
		t.Fail()
	}
	if record != "testtest@gmail.com" {
		t.Fail()
	}

}
