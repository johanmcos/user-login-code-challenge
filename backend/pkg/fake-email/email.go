// Package fake_email presents a workable interface for an email handler
// but using a placeholder implementation
package fake_email

import (
	"fmt"
	"strings"
)

type fakeEmailServer struct{}

// GetNewInstance gets a new instance of our server
func GetNewInstance() *fakeEmailServer {
	return &fakeEmailServer{}
}

// Send "sends" an email
func (s *fakeEmailServer) Send(address, msg string) error {
	// check if the email address is valid
	if !strings.Contains(address, "@") {
		// TODO this should use a regex to do a more complex check
		return fmt.Errorf("%s is not a valid email address", address)
	}
	return nil
}
