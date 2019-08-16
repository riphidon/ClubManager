package utils

import (
	"crypto/md5"
	"fmt"
	"net/http"
	"time"

	"github.com/riphidon/clubmanager/config"
)

func SessionCookie(w http.ResponseWriter) {
	data := config.Data.CookieHash + time.Now().Format(config.Data.CookieBlock)
	hashed := md5.Sum([]byte(data))
	expiration := time.Now().Add(365 * 24 * time.Hour)
	http.SetCookie(w, &http.Cookie{Name: "session", Value: fmt.Sprintf("%v", hashed), SameSite: http.SameSiteStrictMode, Expires: expiration})
}

func EndSession(w http.ResponseWriter, r *http.Request, path string) {
	http.SetCookie(w, &http.Cookie{Name: "session", Value: "Deleted"})
	http.Redirect(w, r, path, http.StatusSeeOther)
}

func CheckSessionCookie(c *http.Cookie) bool {
	dataCheck := config.Data.CookieHash + time.Now().Format(config.Data.CookieBlock)
	hashedCheck := md5.Sum([]byte(dataCheck))
	if c.Name == "session" && c.Value == fmt.Sprintf("%v", hashedCheck) {
		return true
	}
	return false
}

func SetSession(w http.ResponseWriter, path string, segment, param string) string {
	SessionCookie(w)
	p := DoURL(path, segment, param)
	return p
}
