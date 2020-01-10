package services

import (
	"net/http"

	"github.com/riphidon/clubmanager/db"
	"github.com/riphidon/clubmanager/utils"
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
	group, err := db.Role(email)
	if err != nil {
		return c, err
	}
	c.ID = user
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

func RedirectOnLogin(email string, w http.ResponseWriter, r *http.Request) error {
	cred, err := UserCredentials(email)
	if err != nil {
		return err
	}
	group := cred.Group
	id := cred.ID
	utils.SetCookieHandler(w, r, id)
	if group == "user" {
		http.Redirect(w, r, "/profile", http.StatusFound)
		return nil
	}
	http.Redirect(w, r, "/admin", http.StatusFound)
	return nil
}
