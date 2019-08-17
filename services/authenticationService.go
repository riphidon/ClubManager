package services

import (
	"fmt"
	"net/http"

	"github.com/riphidon/clubmanager/db"
	"golang.org/x/crypto/bcrypt"
)

type credentials struct {
	ID    string
	Group string
}

func UserCredentials(email string) (credentials, error) {
	var c credentials
	user, err := db.ID(email)
	if err != nil {
		return c, err
	}
	fmt.Printf("user : %v\n", user)
	group, err := db.Role(email)
	if err != nil {
		return c, err
	}
	fmt.Printf("group : %v\n", group)
	c.ID = user
	fmt.Printf("c.ID: %v\n", c.ID)
	c.Group = group
	return c, nil
}

func ValidatePassword(email, hash string) bool {
	h, err := db.Hash(email)
	if err != nil {
		return false
	}
	if err := bcrypt.CompareHashAndPassword([]byte(h), []byte(hash)); err != nil {
		return false
	}
	return true
}

func RedirectByGroup(email string, w http.ResponseWriter, r *http.Request) error {
	cred, err := UserCredentials(email)
	if err != nil {
		return err
	}

	group := cred.Group
	if group == "users" {
		http.Redirect(w, r, "/profile", http.StatusFound)
		return nil
	}
	http.Redirect(w, r, "/admin", http.StatusFound)
	return nil
}
