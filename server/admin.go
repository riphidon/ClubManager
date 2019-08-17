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
	list := interfaces.ListAllUsers()
	beltList := interfaces.ViewBeltList()
	adminPage := utils.SplitPath(r.URL.Path, 1)
	switch adminPage {
	case "profiles":
		//switchCaseProfiles(w, r, userData)
		var paramList []*models.ClubUser
		q := r.URL.Query()
		switch q.Get("do") {
		case "search":
			fmt.Println("INSIDE SEARCH")
			r.ParseForm()
			searchBy := r.FormValue("userSelector")
			fmt.Printf("searchBy: %v\n", searchBy)
			belt := r.FormValue("belt")
			if searchBy == "rank" {
				paramList = interfaces.FindUserByRank(belt)
			}

			err := RenderPage(w, config.Data.AdminPath, "profiles", &Page{UserList: list, User: userData, BeltList: beltList, ListByParam: paramList})
			if err != nil {
				return err
			}

		default:
			err := RenderPage(w, config.Data.AdminPath, "profiles", &Page{Title: "profiles", UserList: list, User: userData, BeltList: beltList, ListByParam: paramList})
			if err != nil {
				return err
			}

		}
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
		err = RenderPage(w, config.Data.AdminPath, "admin", &Page{User: userData})
		if err != nil {
			return err
		}
	}

	return nil
}

func switchCaseProfiles(w http.ResponseWriter, r *http.Request, userData models.ClubUser) error {
	list := interfaces.ListAllUsers()
	beltList := interfaces.ViewBeltList()
	var paramList []*models.ClubUser
	q := r.URL.Query()
	switch q.Get("do") {
	case "search":
		fmt.Println("INSIDE SEARCH")
		r.ParseForm()
		searchBy := r.FormValue("userSelector")
		fmt.Printf("searchBy: %v\n", searchBy)
		belt := r.FormValue("belt")
		if searchBy == "rank" {
			paramList = interfaces.FindUserByRank(belt)
		}

		err := RenderPage(w, config.Data.AdminPath, "profiles", &Page{UserList: list, User: userData, BeltList: beltList, ListByParam: paramList})
		if err != nil {
			return err
		}

	default:
		err := RenderPage(w, config.Data.AdminPath, "profiles", &Page{Title: "profiles", User: userData})
		if err != nil {
			return err
		}

	}
	err := RenderPage(w, config.Data.AdminPath, "profiles", &Page{Title: "profiles", User: userData})
	if err != nil {
		return err
	}
	return nil
}
