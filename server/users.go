package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/riphidon/clubmanager/config"
	"github.com/riphidon/clubmanager/db"
	"github.com/riphidon/clubmanager/interfaces"
	"github.com/riphidon/clubmanager/services"
	"github.com/riphidon/clubmanager/utils"
)

var message string = ""

//Users SECTION
func Home(w http.ResponseWriter, r *http.Request) error {
	var err error
	infos, err := db.GetInfos()
	if err != nil {
		return err
	}
	events, err := db.GetEvents()
	if err != nil {
		return err
	}
	state := utils.CheckState(id)
	err = RenderPage(w, config.Data.UserPath, "home", &Page{Title: "home", LogState: state, Infos: infos, Events: events})
	if err != nil {
		return err
	}
	return nil
}

func RegisterPage(w http.ResponseWriter, r *http.Request) error {
	userErr := ""
	beltList := interfaces.ViewBeltList()
	state := false
	if r.Method == "POST" {
		r.ParseForm()
		email := r.FormValue("email")
		errCode, err := interfaces.Register(w, r, email)
		if err != nil {
			return err
		}
		//handle error display for client
		if errCode != 0 {
			userErr := utils.CatchUserErr(errCode)
			fmt.Printf("userErr: %v", userErr)
			err := RenderPage(w, config.Data.UserPath, "register", &Page{Title: "register", BeltList: beltList, UserErr: userErr, LogState: state})
			if err != nil {
				return err
			}
			return nil
		}
		services.RedirectOnLogin(email, w, r)
	}
	err := RenderPage(w, config.Data.UserPath, "register", &Page{Title: "register", BeltList: beltList, UserErr: userErr})
	if err != nil {
		return err
	}
	return nil
}

func LoginPage(w http.ResponseWriter, r *http.Request) error {
	state := false
	q := r.URL.Query()
	switch q.Get("do") {
	case "logout":
		id = ""
		utils.EndSession(w, r, "/")
		//http.Redirect(w, r, "/", http.StatusFound)
		return nil
	case "login":
		if r.Method == "POST" {
			fmt.Println("inside login")
			r.ParseForm()
			email := r.FormValue("identifier")
			code, err := interfaces.Login(w, r, email)
			if err != nil {
				return err
			}
			//handle error display for client
			if code != 0 {
				userErr := utils.CatchUserErr(code)
				fmt.Printf("userErr: %v\n", userErr)
				err := RenderPage(w, config.Data.UserPath, "login", &Page{Title: "login", UserErr: userErr, LogState: state})
				if err != nil {
					return err
				}
				return nil
			}
			if err := services.RedirectOnLogin(email, w, r); err != nil {
				fmt.Printf("err: %v\n", err)
				return err
			}
			return nil
		}
	default:
		err := RenderPage(w, config.Data.UserPath, "login", &Page{Title: "login", Message: message})
		return err
	}

	return nil
}

func ProfilePage(w http.ResponseWriter, r *http.Request) error {
	var err error
	state := utils.CheckState(id)
	fmt.Printf("id profilepage is %v\n", id)
	userData := interfaces.ViewUserData(r, id, "user")
	q := r.URL.Query()
	switch q.Get("do") {
	case "edit":
		beltList := interfaces.ViewBeltList()
		userData := interfaces.ViewUserData(r, id, "user")
		if err := RenderPage(w, config.Data.UserPath, "editProfile", &Page{Title: "edition", User: userData, BeltList: beltList, LogState: state}); err != nil {
			return err
		}
	case "upd":
		if r.Method == "POST" {
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
		err = RenderPage(w, config.Data.UserPath, "profile", &Page{Title: "profile", User: userData, LogState: state})
		if err != nil {
			return err
		}
	}
	return nil
}

func CompRegisteration(w http.ResponseWriter, r *http.Request) error {
	fmt.Printf("message: %v\n", message)
	state := utils.CheckState(id)
	if state == false {
		message = "Veuillez vous connecter afin de vous inscrire à cet évènement"
		fmt.Printf("message: %v\n", message)
		http.Redirect(w, r, "login", http.StatusFound)
	} else {
		userData := interfaces.ViewUserData(r, id, "user")
		weights := interfaces.ViewWeightList()
		ages := interfaces.ViewAgeList()
		eventTitle := r.URL.Query().Get("event")
		fmt.Printf("log: %v\n", state)
		err := RenderPage(w, config.Data.UserPath, "compRegisteration", &Page{Title: "competition", User: userData, WeightList: weights, AgeList: ages, EventTitle: eventTitle, LogState: state})
		if err != nil {
			return err
		}
	}

	return nil
}
