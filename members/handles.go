//Package members is for member CRUD
package members

import (
	"encoding/json"
	_ "github.com/lib/pq" // github.com/lib/pq
	"github.com/gorilla/mux"
	"github.com/maxdobeck/basic-app-api/models"
	"github.com/maxdobeck/basic-app-api/sessions"
	"log"
	"net/http"
)

type memberOutput struct {
	Status string
	Errors []string
}

type resDetails struct {
	Status  string
	Message []string
}

type member struct {
	NewName, NewEmail1, NewEmail2 string
}

// SignupMember creates a single member
func SignupMember(w http.ResponseWriter, r *http.Request) {
	var memberValid = true
	var m models.NewMember
	var signupErrs []string
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		log.Println("Error decoding new member >>", err)
		log.Println("Bad New Member data provided: ", r.Body)
		signupErrs = append(signupErrs, "Error processing new member data.")
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

// UpdateMemberEmail allows the user to update member information and returns an error or the newly made member name
func UpdateMemberEmail(w http.ResponseWriter, r *http.Request) {
	if sessions.GoodSession(r) != true {
		msg := resDetails{
			Status:  "Expired session or cookie",
			Message: []string{"Session Expired.  Log out and log back in."},
		}
		json.NewEncoder(w).Encode(msg)
		return
	}

	var msg resDetails
	vars := mux.Vars(r)
	if vars["id"] == "" {
		log.Println("Unexpected URL:", r.URL)
		msg.Status = "Error"
		msg.Message = append(msg.Message, "Path is unexpected.")
		json.NewEncoder(w).Encode(msg)
		return
	}

	var memberUpdate member
	err := json.NewDecoder(r.Body).Decode(&memberUpdate)
	if err != nil {
		log.Println("Error decoding body >>", err)
	}
	// Check for bad email length
	/* if len(memberUpdate.NewEmail1) < 5 || len(memberUpdate.NewEmail2) < 5 {
		msg := resDetails{
			Status:  "Bad Name",
			Message: append(msg.Message, "Email must have more than 0 characters."),
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msg)
	} */
	log.Println("New Email: ", memberUpdate.NewEmail1)
	if len(memberUpdate.NewEmail1) >= 5 || len(memberUpdate.NewEmail2) >= 5 {
		if memberUpdate.NewEmail1 != memberUpdate.NewEmail2 {
			msg := resDetails{
				Status:  "Bad Email",
				Message: append(msg.Message, "Emails don't match."),
			}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(msg)
			return
		} else if models.UpdateMemberEmail(vars["id"], memberUpdate.NewEmail1) == true {
			msg.Message = append(msg.Message, memberUpdate.NewEmail1)
			msg.Status = "OK"
			json.NewEncoder(w).Encode(msg)
		}
	}

	log.Println("Path Variables: ", vars)
	log.Println("Member's ID: ", vars["id"])
	log.Println(msg)
}

// UpdateMemberName will update the existing member name for authorized sessions
func UpdateMemberName(w http.ResponseWriter, r *http.Request) {
	if sessions.GoodSession(r) != true {
		msg := resDetails{
			Status:  "Expired session or cookie",
			Message: []string{"Session Expired.  Log out and log back in."},
		}
		json.NewEncoder(w).Encode(msg)
		return
	}

	var msg resDetails
	vars := mux.Vars(r)
	if vars == nil || vars["id"] == "" {
		msg.Status = "Error"
		msg.Message = append(msg.Message, "Path is unexpected.")
		json.NewEncoder(w).Encode(msg)
		return
	}

	var memberUpdate member
	err := json.NewDecoder(r.Body).Decode(&memberUpdate)
	if err != nil {
		log.Println("Error decoding body >>", err)
	}
	// Check for bad name length
	/*if len(memberUpdate.NewName) < 1 {
		msg := resDetails{
			Status:  "Bad Name",
			Message: append(msg.Message, "Name must have more than 0 characters."),
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msg)
	} */
	log.Println("New Name: ", memberUpdate.NewName)
	if len(memberUpdate.NewName) >= 1 {
		if models.UpdateMemberName(vars["id"], memberUpdate.NewName) == true {
			msg.Message = append(msg.Message, memberUpdate.NewName)
			msg.Status = "OK"
			json.NewEncoder(w).Encode(msg)
		}
	}

	log.Println("Path Variables: ", vars)
	log.Println("Member's ID: ", vars["id"])
	log.Println(msg)
}
