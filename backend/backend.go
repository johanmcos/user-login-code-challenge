package login_backend

import (
	"github.com/johanmcos/user-login-code-challenge/backend/pkg/database"
	fake_email "github.com/johanmcos/user-login-code-challenge/backend/pkg/fake-email"
	"net/http"
)

// EmailServer defines a generic interface for an email server
// allows potentially using a real one in the future
type EmailServer interface {
	Send(recipient, msg string) error
}

// User is an individual user
type User interface {
	VerifyPassword(password string) bool
}

// Database defines a generic database
type Database interface {
	GetUser(username string) *User
	AddUser(*User) error
}

// holds values that persist between requests
type backend struct {
	mailServer *EmailServer
	db         *Database
}

// register our login handler and then listen for incoming requests
func main() {
	var mailServer EmailServer = fake_email.GetNewInstance()
	var db Database = database.CreateDatabase()
	backend := &backend{
		&mailServer,
		&db,
	}
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/forgot", forgotHandler)
	http.ListenAndServe("localhost:8080", nil)
}

func loginHandler(resp http.ResponseWriter, req *http.Request) {
	return
}

func registerHandler(resp http.ResponseWriter, req *http.Request) {
	return
}

func forgotHandler(resp http.ResponseWriter, req *http.Request) {
	return
}
