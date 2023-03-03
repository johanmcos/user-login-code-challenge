package main

import (
	"github.com/johanmcos/user-login-code-challenge/backend/pkg/database"
	"github.com/johanmcos/user-login-code-challenge/backend/pkg/user"
	"testing"
)

func TestDatabase(t *testing.T) {
	db := database.CreateDatabase()
	if nobody, err := db.GetUser("johndoe"); nobody != nil || err == nil {
		t.Error("unexpected behavior when looking for nonexistent nobody")
	}
	testUser := user.NewUser("johndoe", "password")
	err := db.AddUser(testUser)
	if err != nil {
		t.Errorf("error when adding user: %v", err)
	}
	johnDoe, err := db.GetUser("johndoe")
	if err != nil {
		t.Errorf("error when retrieving user: %v", err)
	}
	if johnDoe.Name != "johndoe" {
		t.Errorf("retrieved user does not have the expected traits %.v", johnDoe)
	}
}

func TestUser(t *testing.T) {
	user1 := user.NewUser("user1", "password")
	user2 := user.NewUser("user2", "password")
	if user1.HashedPassword == user2.HashedPassword {
		t.Error("Identical passwords lead to identical hashes, salt fail")
	}
	if works := user1.VerifyPassword("notpassword"); works {
		t.Error("logging in with wrong password works")
	}
}
