package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/riphidon/clubmanager/config"
	"github.com/riphidon/clubmanager/utils"
)

// type Handler struct {
// 	logger *log.Logger
// }
// func (h *Handler) Logger(next http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		h.logger.Printf("Logged, %v requested %v", r.RemoteAddr, r.URL)
// 		next(w, r)
// 	}
// }
// func NewHandler(log *log.Logger) *Handler {
// 	return &Handler{
// 		logger: log,
// 	}
// }

// func (fn AuthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	log.Printf("logged %v requested %v", r.RemoteAddr, r.URL)
// 	cookie, errCookie := r.Cookie("session")
// 	if errCookie != nil || utils.CheckSessionCookie(cookie) == false {
// 		http.Redirect(w, r, "/login", http.StatusFound)
// 		return
// 	}
// 	fmt.Printf("checkSession : %v\n", utils.CheckSessionCookie(cookie))
// 	err := fn(w, r)
// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("error : %+v\n", err), http.StatusInternalServerError)
// 	}
// }

// type auth struct {
// 	fn func(w http.ResponseWriter, r *http.Request)
// }

type AppHandler func(w http.ResponseWriter, r *http.Request) error

type AuthHandler func(w http.ResponseWriter, r *http.Request) error

func (fn AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("logged %v requested %v", r.RemoteAddr, r.URL)
	err := fn(w, r)
	if err != nil {
		http.Error(w, fmt.Sprintf("error : %+v\n", err), http.StatusInternalServerError)
	}
}

func (fn AuthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("logged %v requested %v", r.RemoteAddr, r.URL)
	utils.SetCookieHandler(w, r)
	cookie, errCookie := r.Cookie("session")
	if errCookie != nil {
		return
	}
	fmt.Printf("cookie value is : %v\n", cookie.Value)
	value, err := utils.ReadCookieHandler(w, r)
	if err != nil {
		fmt.Printf("error reading cookie value: %v\n", err)
	}
	fmt.Printf("value is : %v\n", value)
	if errCookie != nil || utils.CheckCookie(cookie, value) == false {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	fmt.Printf("checkSession : %v\n", utils.CheckCookie(cookie, value))
	err = fn(w, r)
	if err != nil {
		http.Error(w, fmt.Sprintf("error : %+v\n", err), http.StatusInternalServerError)
	}
}

//midleware for routes management
func SetupRoutes(mux *http.ServeMux) {
	mux.Handle("/static/", config.FileSystem)
	mux.Handle("/", AppHandler(Home))
	mux.Handle("/register", AppHandler(RegisterPage))
	mux.Handle("/login", AppHandler(LoginPage))
	mux.Handle("/profile/", AuthHandler(ProfilePage))
	mux.Handle("/admin/", AuthHandler(AdminSection))

}

// mux.Handle("/", routes.Home)
// mux.Handle("/login", routes.Login)
// mux.Handle("/register", routes.Register)
// mux.Handle("/profile/", routes.Profile)
