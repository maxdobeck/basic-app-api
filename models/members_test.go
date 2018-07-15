package models

import (
	"log"
	"os"
	"testing"
)

func TestCreateMember(t *testing.T) {
	ConnToDB(os.Getenv("PGURL"))

	_, delErr := Db.Query("DELETE FROM members WHERE email like 'testtest@gmail.com'")
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
	err := Db.QueryRow("SELECT email FROM members WHERE email like 'testtest@gmail.com'").Scan(&record)
	if err != nil {
		log.Println(err)
		t.Log(err)
		t.Fail()
	}
	if record != "testtest@gmail.com" {
		t.Fail()
	}

}
