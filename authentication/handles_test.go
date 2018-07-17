package authentication

import (
	"fmt"
	"github.com/maxdobeck/basic-app-api/members"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// An HTTP test to ensure a login request is rejected if the credentials are wrong
func TestLoginGoodCredentials(t *testing.T) {
	// Signup a user
	signupBody := strings.NewReader(`{"email": "goodUser@gmail.com", "email2":"goodUser@gmail.com", "password": "supersecret", "password2":"supersecret", "name":"Valid User Signup"}`)
	signupReq, signupErr := http.NewRequest("POST", "/members", signupBody)
	if signupErr != nil {
		t.Fail()
	}
	wSignup := httptest.NewRecorder()
	members.SignupMember(wSignup, signupReq)

	bodyReader := strings.NewReader(`{"email": "WrongEmail@email.com", "password": "wrongPassword"}`)

	req, err := http.NewRequest("POST", "/login", bodyReader)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	Login(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 401 {
		t.Fail()
	}

	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Println(string(body))
}

// Test the Login command with a valid set of credentials
func TestLoginBadCredentials(t *testing.T) {
	// Signup a user
	signupBody := strings.NewReader(`{"email": "testValidCreds@gmail.com", "email2":"testValidCreds@gmail.com", "password": "supersecret", "password2":"supersecret", "name":"Valid User Signup"}`)
	signupReq, signupErr := http.NewRequest("POST", "/members", signupBody)
	if signupErr != nil {
		t.Fail()
	}
	wSignup := httptest.NewRecorder()
	members.SignupMember(wSignup, signupReq)

	bodyReader := strings.NewReader(`{"email": "testValidCreds@gmail.com", "password": "supersecret"}`)
	req, err := http.NewRequest("POST", "/login", bodyReader)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	Login(w, req)

	resp := w.Result()

	if resp.StatusCode != 200 {
		t.Fail()
	}
}
