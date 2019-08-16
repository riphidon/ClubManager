package models

import (
	"crypto/md5"
	"fmt"
	"net/http"
	"time"

	"github.com/riphidon/clubmanager/config"

	"golang.org/x/crypto/bcrypt"
)

type Member struct {
	User_id      int
	FirstName    string
	LastName     string
	Email        string
	Password     string
	Rank         string
	Licence      string
	MedicalCert  bool
	Competitor   bool
	Category     string
	RegisterDate time.Time
}

func (m Member) ParseNewMemberData(w http.ResponseWriter, r *http.Request) (Member, error) {
	var err error
	r.ParseForm()
	m.FirstName = r.FormValue("firstname")
	m.LastName = r.FormValue("lastname")
	m.Email = r.FormValue("email")
	m.Password, err = HashPassword(r.FormValue("password"))
	if err != nil {
		return m, err
	}
	m.Rank = r.FormValue("belt")

	fmt.Printf("hashed: %v", m.Password)
	return m, nil
}

func HashPassword(password string) (string, error) {
	bytePass := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(bytePass, 7)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func SessionCookie(w http.ResponseWriter) {
	data := config.Data.CookieHash + time.Now().Format(config.Data.CookieBlock)
	hashed := md5.Sum([]byte(data))
	expiration := time.Now().Add(365 * 24 * time.Hour)
	http.SetCookie(w, &http.Cookie{Name: "session", Value: fmt.Sprintf("%v", hashed), SameSite: http.SameSiteStrictMode, Expires: expiration})
}

func CheckSessionCookie(c *http.Cookie) bool {
	dataCheck := config.Data.CookieHash + time.Now().Format(config.Data.CookieBlock)
	hashedCheck := md5.Sum([]byte(dataCheck))
	if c.Name == "session" && c.Value == fmt.Sprintf("%v", hashedCheck) {
		return true
	}
	return false
}
