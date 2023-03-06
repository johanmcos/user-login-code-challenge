package main

import (
	"encoding/json"
	"fmt"
	db "github.com/johanmcos/user-login-code-challenge/backend/pkg/database"
	"github.com/johanmcos/user-login-code-challenge/backend/pkg/user"
	"io"
	"net/http"
)

// holds values that persist between requests
type backend struct {
	db *db.Database
}

// expected request from client
type Request struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// register our login handler and then listen for incoming requests
func main() {
	b := &backend{
		db.CreateDatabase(),
	}
	http.HandleFunc("/login", b.loginHandler)
	http.HandleFunc("/register", b.registerHandler)
	http.HandleFunc("/hello", func(w http.ResponseWriter, _ *http.Request) {
		_, _ = io.WriteString(w, "Hello to you as well")
	})
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		panic(err)
	}
}

func (b *backend) loginHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "text/plain")
	username, password, ok := parseRequest(req)
	if !ok {
		w.WriteHeader(http.StatusForbidden)
		_, err := io.WriteString(w, "basic auth missing or malformed")
		if err != nil {
			fmt.Println(err)
		}
		return
	}
	thisUser, err := b.db.GetUser(username)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		_, err := io.WriteString(w, "username/password not valid")
		if err != nil {
			fmt.Println(err)
		}
		return
	}
	if thisUser.VerifyPassword(password) {
		_, err := io.WriteString(w, "login successful")
		if err != nil {
			fmt.Println(err)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		_, err := io.WriteString(w, "username/password not valid")
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (b *backend) registerHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "text/plain")
	username, password, ok := parseRequest(req)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		_, err := io.WriteString(w, "basic auth missing or malformed")
		if err != nil {
			fmt.Println(err)
		}
		return
	}
	if b.db.UserExists(username) {
		w.WriteHeader(http.StatusConflict)
		_, err := io.WriteString(w, "username must be unique")
		if err != nil {
			fmt.Println(err)
		}
		return
	}
	newUser := user.NewUser(username, password)
	err := b.db.AddUser(newUser)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		_, err := io.WriteString(w, "error when attempting to create new user")
		if err != nil {
			fmt.Println(err)
		}
		return
	}
	_, _ = io.WriteString(w, "Registration successful!")
}

func parseRequest(req *http.Request) (username, password string, ok bool) {
	userRequest := &Request{}
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, userRequest)
	if err != nil {
		return
	}
	return userRequest.Username, userRequest.Password, true
}
