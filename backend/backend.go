package main

import (
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
	username, password, ok := req.BasicAuth()
	if !ok {
		_, err := io.WriteString(w, "basic auth missing or malformed")
		if err != nil {
			fmt.Println(err)
		}
		w.WriteHeader(http.StatusForbidden)
		return
	}
	thisUser, err := b.db.GetUser(username)
	if err != nil {
		_, err := io.WriteString(w, "username/password not valid")
		if err != nil {
			fmt.Println(err)
		}
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if thisUser.VerifyPassword(password) {
		_, err := io.WriteString(w, "login successfull")
		if err != nil {
			fmt.Println(err)
		}
		w.WriteHeader(http.StatusOK)
	} else {
		_, err := io.WriteString(w, "username/password not valid")
		if err != nil {
			fmt.Println(err)
		}
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func (b *backend) registerHandler(w http.ResponseWriter, req *http.Request) {
	username, password, ok := req.BasicAuth()
	if !ok {
		_, err := io.WriteString(w, "basic auth missing or malformed")
		if err != nil {
			fmt.Println(err)
		}
		w.WriteHeader(http.StatusBadRequest)
	}
	newUser := user.NewUser(username, password)
	err := b.db.AddUser(newUser)
	if err != nil {
		fmt.Println(err)
		_, err := io.WriteString(w, "error when attempting to create new user")
		if err != nil {
			fmt.Println(err)
		}
		w.WriteHeader(http.StatusBadRequest)
	}
}
