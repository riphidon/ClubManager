package utils

import (
	"fmt"
	"net/http"

	"github.com/gorilla/securecookie"
)

var hashKey = []byte(securecookie.GenerateRandomKey(32))
var blockKey = []byte(securecookie.GenerateRandomKey(32))
var s = securecookie.New(hashKey, blockKey)

// func SessionCookie(w http.ResponseWriter) {
// 	data := config.Data.CookieHash + time.Now().Format(config.Data.CookieBlock)
// 	hashed := md5.Sum([]byte(data))
// 	expiration := time.Now().Add(365 * 24 * time.Hour)
// 	http.SetCookie(w, &http.Cookie{Name: "session", Value: fmt.Sprintf("%v", hashed), SameSite: http.SameSiteStrictMode, Expires: expiration})
// }

func EndSession(w http.ResponseWriter, r *http.Request, path string) {
	http.SetCookie(w, &http.Cookie{Name: "session", Value: "Deleted"})
	http.Redirect(w, r, path, http.StatusSeeOther)
}

// func CheckSessionCookie(c *http.Cookie) bool {
// 	dataCheck := config.Data.CookieHash + time.Now().Format(config.Data.CookieBlock)
// 	hashedCheck := md5.Sum([]byte(dataCheck))
// 	if c.Name == "session" && c.Value == fmt.Sprintf("%v", hashedCheck) {
// 		return true
// 	}
// 	return false
// }

// func SetSession(w http.ResponseWriter, path string, segment, param string) string {
// 	SessionCookie(w)
// 	p := DoURL(path, segment, param)
// 	return p
// }

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
	fmt.Printf("The value of id is %q", value["id"])
	id := value["id"]
	return id, nil
}

func CheckCookie(c *http.Cookie, value string) bool {
	fmt.Printf("c.Name: %v, c.Value: %v, value :%v", c.Name, c.Value, value)
	if c.Name == "session" && c.Value == value {
		return true
	}
	return false
}
