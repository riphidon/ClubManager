package interfaces

import (
	"fmt"
	"net/http"

	"github.com/riphidon/clubmanager/services"
)

func Register(w http.ResponseWriter, r *http.Request, email string) (int, error) {
	code := 0
	r.ParseForm()
	c := r.FormValue("passphrase")
	if services.PassPhrase(c) == false {
		code = 1
		return code, nil
	} else {
		name := r.FormValue("lastname")
		firstname := r.FormValue("firstname")
		email := email
		belt := r.FormValue("belt")
		existence, err := services.CheckUserExists(email)
		if err != nil {
			return code, err
		}
		fmt.Printf("existence : %v\n", existence)
		if existence != false {
			code = 2
			return code, nil
		}
		password, err := services.HashPassword(r.FormValue("password"))
		if err != nil {
			return code, err
		}
		n := services.NewUser(name, firstname, email, password, belt)
		if err := services.AddNewUser(n); err != nil {
			return code, err
		}
	}
	return code, nil
}

func Login(w http.ResponseWriter, r *http.Request, email string) (int, error) {
	code := 0
	r.ParseForm()
	pass := r.FormValue("password")
	isUserValid, _ := services.CheckUserExists(email)
	if isUserValid != true {
		code = 3
		return code, nil
	}
	isPassValid := services.ValidatePassword(email, pass)
	if isPassValid != true {
		code = 4
		return code, nil
	}
	return code, nil
}
