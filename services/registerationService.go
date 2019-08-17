package services

import (
	"time"

	"github.com/riphidon/clubmanager/db"
	"github.com/riphidon/clubmanager/models"
	"golang.org/x/crypto/bcrypt"
)

func NewUser(name, firstname, email, password, belt string) models.ClubUser {
	return models.ClubUser{
		Name:      name,
		Firstname: firstname,
		Email:     email,
		Hash:      password,
		Rank:      belt,
		Role:      "user",
		MedCert:   false,
		CreatedOn: time.Now().AddDate(0, 0, 0),
	}
}

func CheckUserExists(email string) (bool, error) {
	exists, err := db.UserExists(email)
	if err != nil {
		return exists, err
	}
	return exists, nil
}

func HashPassword(password string) (string, error) {
	bytePass := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(bytePass, 7)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func PassPhrase(key string) bool {
	pass := "selva life"
	if key != pass {
		return false
	}
	return true
}

func AddNewUser(u models.ClubUser) error {
	if err := db.StoreNewUser(u); err != nil {
		return err
	}
	return nil
}
