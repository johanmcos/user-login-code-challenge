package login_backend

import (
	fake_email "github.com/johanmcos/user-login-code-challenge/backend/pkg/fake-email"
	"net/http"
)

type EmailServer interface {
	Send(recipient, msg string) error
}

var mailServer EmailServer

// register our login handler and then listen for incoming requests
func main() {
	mailServer = fake_email.GetNewInstance()
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
