package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/riphidon/clubmanager/config"
	"github.com/riphidon/clubmanager/interfaces"
	"github.com/riphidon/clubmanager/services"
	"github.com/riphidon/clubmanager/utils"
)

//Users SECTION
func Home(w http.ResponseWriter, r *http.Request) error {
	err := RenderPage(w, config.Data.UserPath, "home", &Page{Title: "home"})
	if err != nil {
		return err
	}
	return nil
}

func RegisterPage(w http.ResponseWriter, r *http.Request) error {
	userErr := ""
	beltList := interfaces.ViewBeltList()
	if r.Method == "POST" {
		r.ParseForm()
		email := r.FormValue("email")
		errCode, err := interfaces.Register(w, r, email)
		if err != nil {
			return err
		}
		if errCode != 0 {
			userErr := utils.CatchUserErr(errCode)
			fmt.Printf("userErr: %v", userErr)
			err := RenderPage(w, config.Data.UserPath, "register", &Page{Title: "register", BeltList: beltList, UserErr: userErr})
			if err != nil {
				return err
			}
			return nil
		}
		services.RedirectOnRegister(email, w, r)
	}
	err := RenderPage(w, config.Data.UserPath, "register", &Page{Title: "register", BeltList: beltList, UserErr: userErr})
	if err != nil {
		return err
	}
	return nil
}

func LoginPage(w http.ResponseWriter, r *http.Request) error {

	q := r.URL.Query()
	switch q.Get("do") {
	case "logout":
		utils.EndSession(w, r, "/")
		http.Redirect(w, r, "/login", http.StatusFound)
		return nil
	case "login":
		if r.Method == "POST" {
			r.ParseForm()
			email := r.FormValue("identifier")
			code, err := interfaces.Login(w, r, email)
			if err != nil {
				return err
			}
			if code != 0 {
				userErr := utils.CatchUserErr(code)
				fmt.Printf("userErr: %v", userErr)
				err := RenderPage(w, config.Data.UserPath, "login", &Page{Title: "login", UserErr: userErr})
				if err != nil {
					return err
				}
				return nil
			}
			if err := services.RedirectByGroup(email, w, r); err != nil {
				return err
			}
			return nil
		}

	default:
		err := RenderPage(w, config.Data.UserPath, "login", &Page{Title: "login"})
		return err
	}

	err := RenderPage(w, config.Data.UserPath, "login", &Page{Title: "login"})
	if err != nil {
		return err
	}
	return nil
}

func ProfilePage(w http.ResponseWriter, r *http.Request) error {
	var err error
	q := r.URL.Query()
	switch q.Get("do") {
	case "edit":
		fmt.Println("inside edit")
		beltList := interfaces.ViewBeltList()
		id := utils.CatchURLData(r, "q")
		userData := interfaces.ViewUserData(r, id, "user")
		fmt.Printf("user data : %v\n", userData)
		if err := RenderPage(w, config.Data.UserPath, "editProfile", &Page{Title: "edition", User: userData, BeltList: beltList}); err != nil {
			return err
		}
	case "upd":
		if r.Method == "POST" {
			fmt.Println("inside post")
			member := r.URL.Query().Get("q")
			user, err := strconv.Atoi(member)
			if err != nil {
				return err
			}
			if err := interfaces.Edit(w, r, user); err != nil {
				return err
			}
			http.Redirect(w, r, "/profile", http.StatusFound)
		}
	default:
		id := utils.CatchURLData(r, "q")
		userData := interfaces.ViewUserData(r, id, "user")
		err = RenderPage(w, config.Data.UserPath, "profile", &Page{Title: "profile", User: userData})
		if err != nil {
			return err
		}
	}
	return nil
}
