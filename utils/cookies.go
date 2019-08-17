package utils

import (
	"fmt"
	"net/http"

	"github.com/gorilla/securecookie"
)

var hashKey = []byte(securecookie.GenerateRandomKey(32))
var blockKey = []byte(securecookie.GenerateRandomKey(32))
var s = securecookie.New(hashKey, blockKey)

func EndSession(w http.ResponseWriter, r *http.Request, path string) {
	http.SetCookie(w, &http.Cookie{Name: "session", Value: "Deleted"})
	http.Redirect(w, r, path, http.StatusSeeOther)
}

func SetCookieHandler(w http.ResponseWriter, r *http.Request, id string) {
	value := map[string]string{
		"id": id,
	}
	encoded, err := s.Encode("session", value)
	if err == nil {
		cookie := &http.Cookie{
			Name:  "session",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(w, cookie)
	}
	fmt.Printf("error in SetCookie: %v\n", err)
}

func ReadCookieHandler(w http.ResponseWriter, r *http.Request) (string, error) {
	cookie, err := r.Cookie("session")
	if err != nil {
		fmt.Printf("error Decode : %v\n", err)
		return "", err
	}
	value := make(map[string]string)
	if err := s.Decode("session", cookie.Value, &value); err != nil {
		return "", err
	}
	id := value["id"]
	return id, nil
}

func CheckCookie(c *http.Cookie, value string) bool {
	if c.Name == "session" && c.Value == value {
		return true
	}
	return false
}
