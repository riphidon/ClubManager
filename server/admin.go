package server

import (
	"fmt"
	"net/http"

	"github.com/riphidon/clubmanager/config"
	"github.com/riphidon/clubmanager/interfaces"
	"github.com/riphidon/clubmanager/models"
	"github.com/riphidon/clubmanager/utils"
)

//Admin SECTION
func AdminSection(w http.ResponseWriter, r *http.Request) error {
	var err error
	id := Id
	userData := interfaces.ViewUserData(r, id, "admin")
	adminPage := utils.SplitPath(r.URL.Path, 1)
	switch adminPage {
	case "profiles":
		err = switchCaseProfiles(w, r, userData)
	case "tasks":
		userData := interfaces.ViewUserData(r, id, "admin")
		err := RenderPage(w, config.Data.AdminPath, "tasks", &Page{Title: "tasks", User: userData})
		if err != nil {
			return err
		}
	case "competitions":
		err = RenderPage(w, config.Data.AdminPath, "competitions", &Page{Title: "competitions", User: userData})
		if err != nil {
			return err
		}
	default:
		err = RenderPage(w, config.Data.AdminPath, "admin", &Page{Title: "admin", User: userData})
		if err != nil {
			return err
		}
	}

	return nil
}

func switchCaseProfiles(w http.ResponseWriter, r *http.Request, userData models.ClubUser) error {
	q := r.URL.Query()
	switch q.Get("do") {
	case "search":
		r.ParseForm()
		search := r.FormValue("userSelector")
		param := r.FormValue("inputValue")
		fmt.Printf("input: %v\n", param)
		if search != "all" {
			SearchUsers(param, w)
			return nil
		}
		list := interfaces.ListAllUsers()
		err := RenderPage(w, config.Data.AdminPath, "profiles", &Page{UserList: list, User: userData})
		if err != nil {
			return err
		}
		return nil
	default:
		err := RenderPage(w, config.Data.AdminPath, "profiles", &Page{Title: "profiles", User: userData})
		if err != nil {
			return err
		}
		return nil
	}
}

func SearchUsers(param string, w http.ResponseWriter) error {
	switch param {
	case "name":
		listByName := interfaces.FindUserByName(param)
		err := RenderPage(w, config.Data.AdminPath, "profiles", &Page{UserList: listByName})
		if err != nil {
			return err
		}
		return nil
	case "rank":
		listByRank := interfaces.FindUserByName(param)
		err := RenderPage(w, config.Data.AdminPath, "profiles", &Page{UserList: listByRank})
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}
