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

// type AuthHandler struct {
// 	Fn func(w http.ResponseWriter, r *http.Request) error
// 	Id string
// }

var id string

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
	cookie, errCookie := r.Cookie("session")
	if errCookie != nil {
		fmt.Printf("ErrCookie: %v\n", errCookie)
		return
	}
	fmt.Printf("cookie value is : %v\n", cookie.Value)
	value, err := utils.ReadCookieHandler(w, r)
	if err != nil {
		fmt.Printf("error reading cookie value: %v\n", err)
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	id = value
	if errCookie != nil || utils.CheckCookie(cookie, cookie.Value) == false {
		fmt.Printf("error checkCookie: %v\n, %v\n", errCookie, utils.CheckCookie(cookie, cookie.Value))
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	err = fn(w, r)
	if err != nil {
		http.Error(w, fmt.Sprintf("error : %+v\n", err), http.StatusInternalServerError)
	}
}

func favicoHandler(w http.ResponseWriter, r *http.Request) error {
	http.Redirect(w, r, "/static/assets/images/logo.ico", http.StatusSeeOther)
	return nil
}

//midleware for routes management
func SetupRoutes(mux *http.ServeMux) {
	mux.Handle("/static/", config.FileSystem)
	mux.Handle("/", AppHandler(Home))
	mux.Handle("/register", AppHandler(RegisterPage))
	mux.Handle("/login", AppHandler(LoginPage))
	mux.Handle("/profile/", AuthHandler(ProfilePage))
	mux.Handle("/admin/", AuthHandler(AdminSection))
	mux.Handle("/compRegisteration", AuthHandler(CompRegisteration))
	mux.Handle("/favicon.ico", AppHandler(favicoHandler))

}

// mux.Handle("/", routes.Home)
// mux.Handle("/login", routes.Login)
// mux.Handle("/register", routes.Register)
// mux.Handle("/profile/", routes.Profile)
