//Package members is for member CRUD
package members

import (
	"encoding/json"
	"github.com/maxdobeck/gatekeeper/models"
	"log"
	"net/http"
)

type memberOutput struct {
	Status string
	Errors []string
}

// SignupMember creates a single member
func SignupMember(w http.ResponseWriter, r *http.Request) {
	var memberValid = true
	var m models.NewMember
	var signupErrs []string
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		log.Println("Error decoding new member >>", err)
	}

	if m.Name == "" {
		signupErrs = append(signupErrs, "Name must not be empty.")
		// json.NewEncoder(w).Encode("Name must not be empty.")
		memberValid = false
	}

	if emailsMatch(m.Email, m.Email2) != true {
		signupErrs = append(signupErrs, "Emails do not match.")
		// json.NewEncoder(w).Encode("Emails do not match.")
		memberValid = false
	}

	if emailAvailable(m.Email) != true {
		signupErrs = append(signupErrs, "Email is already in use.")
		// json.NewEncoder(w).Encode("Email is already in use.")
		memberValid = false
	}

	if passwordsMatch(m.Password, m.Password2) != true {
		signupErrs = append(signupErrs, "Passwords do not match.")
		// json.NewEncoder(w).Encode("Passwords do not match.")
		memberValid = false
	}

	if memberValid == true {
		msg := memberOutput{
			Status: "Member Created",
			Errors: signupErrs,
		}
		models.CreateMember(&m)
		json.NewEncoder(w).Encode(msg)
		log.Println("User Created", m.Email, m.Name)
	} else {
		log.Println("Error creating member.")
		msg := memberOutput{
			Status: "Member Not Created",
			Errors: signupErrs,
		}
		json.NewEncoder(w).Encode(msg)
	}
	log.Println("User data supplied:", m)
}
