// Package fakemail presents a workable interface for an email handler
// but using a placeholder implementation
package fakemail

import (
	"fmt"
	"strings"
)

type MailServer struct{}

// GetNewInstance gets a new instance of our server
func GetNewInstance() *MailServer {
	return &MailServer{}
}

// Send "sends" an email
func (s *MailServer) Send(address, msg string) error {
	// check if the email address is valid
	if !strings.Contains(address, "@") {
		// TODO this should use a regex to do a more complex check
		return fmt.Errorf("%s is not a valid email address", address)
	}
	return nil
}
