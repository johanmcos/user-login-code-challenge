package user

import (
	"crypto/rand"
	"crypto/sha256"
	"io"
)

const saltSize = 8 // bytes

type User struct {
	Name           string
	HashedPassword string
	Salt           string
}

func (u *User) VerifyPassword(password string) bool {
	return getHash(password, u.Salt) == u.HashedPassword
}

// NewUser creates a new user
func NewUser(username, password string) *User {
	salt := getSalt()
	hashedPassword := getHash(password, salt)
	return &User{
		Name:           username,
		HashedPassword: hashedPassword,
		Salt:           salt,
	}
}

// getHash takes a password and a Salt and turns them into
// a single sha256 hash, returned as a string
func getHash(password, salt string) string {
	toHash := password + salt
	hashBytes := sha256.Sum256([]byte(toHash))
	hash := string(hashBytes[:])
	return hash
}

// getSalt generates a new random Salt.
func getSalt() string {
	salt := make([]byte, saltSize)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		//TODO don't panic
		panic(err)
	}
	return string(salt)
}
