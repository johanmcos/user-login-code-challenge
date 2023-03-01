// Package database implements a ephemeral db
// only remains until the server process is terminated
package database

import (
	"fmt"
	"github.com/johanmcos/user-login-code-challenge/backend/pkg/user"
	"sync"
)

type Database struct {
	mu    *sync.RWMutex
	users map[string]*backend.User
}

func (d *Database) AddUser(user *user.User) error {
	// TODO implement validation
	d.mu.Lock()
	defer d.mu.Unlock()
	d.users[user.Name] = user
	return nil
}

func (d *Database) GetUser(username string) (user *user.User, err error) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	user, found := d.users[username]
	if !found {
		return nil, fmt.Errorf("no user exists with username %s", username)
	}
	return user, nil
}

func CreateDatabase() Database {
	return Database{}
}
