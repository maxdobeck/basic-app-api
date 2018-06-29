package authentication

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// An HTTP test to ensure a login request is rejected if the credentials are wrong
func TestLoginInvalidCredentials(t *testing.T) {
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
func TestLoginValidCredentials(t *testing.T) {
	bodyReader := strings.NewReader(`{"email": "test@gmail.com", "password": "supersecret"}`)

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
