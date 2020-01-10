package server

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/riphidon/clubmanager/config"
	"github.com/riphidon/clubmanager/db"
	"github.com/riphidon/clubmanager/interfaces"
	"github.com/riphidon/clubmanager/models"
	"github.com/riphidon/clubmanager/utils"
)

//Admin SECTION
func AdminSection(w http.ResponseWriter, r *http.Request) error {
	var err error
	userData := interfaces.ViewUserData(r, id, "admin")
	list := interfaces.ListAllUsers()
	beltList := interfaces.ViewBeltList()
	adminPage := utils.SplitPath(r.URL.Path, 1)
	switch adminPage {
	case "profiles":
		//switchCaseProfiles(w, r, userData)
		var paramList []*models.ClubUser
		var userProfile models.ClubUser
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
				err := RenderPage(w, config.Data.AdminPath, "profiles", &Page{UserList: list, User: userData, BeltList: beltList, ListByParam: paramList})
				if err != nil {
					return err
				}
				return nil
			}
			if searchBy == "list" {
				sUser := r.FormValue("userRow")
				user, err := strconv.Atoi(sUser)
				if err != nil {
					fmt.Printf("error occured parsing data: %v\n", err)
				}
				userProfile = interfaces.FindUserById(user)
				err = RenderPage(w, config.Data.AdminPath, "profiles", &Page{UserList: list, User: userData, BeltList: beltList, UserProfile: userProfile})
				if err != nil {
					return err
				}
				return nil
			}

		case "edit":
			sUser := r.URL.Query().Get("n")
			user, err := strconv.Atoi(sUser)
			if err != nil {
				return err
			}
			editable := interfaces.FindUserById(user)
			err = RenderPage(w, config.Data.AdminPath, "profilesEdit", &Page{Title: "profiles", UserList: list, User: userData, Editable: editable, BeltList: beltList})
			if err != nil {
				return err
			}
		case "update":
			if r.Method == "POST" {
				sUser := r.URL.Query().Get("n")
				user, err := strconv.Atoi(sUser)
				if err != nil {
					return err
				}
				if err := interfaces.AdminEdit(w, r, user); err != nil {
					fmt.Printf("couldn't update user: %v\n", err)
					return err
				}
				http.Redirect(w, r, "/admin/profiles", http.StatusFound)
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
	case "event":
		err = switchCaseEvent(w, r)
	case "info":
		err = switchCaseInfo(w, r)
	default:
		err = RenderPage(w, config.Data.AdminPath, "admin", &Page{User: userData})
		if err != nil {
			return err
		}
	}
	if err != nil {
		return err
	}
	return nil
}

func switchCaseEvent(w http.ResponseWriter, r *http.Request) error {
	q := r.URL.Query()
	events, err := db.GetEvents()
	if err != nil {
		return err
	}
	switch q.Get("do") {
	case "new":
		if err := RenderPage(w, config.Data.AdminPath, "addEvent", &Page{Title: "event creation"}); err != nil {
			return err
		}
	case "edit":
		seventID := q.Get("event")
		eventID, err := strconv.Atoi(seventID)
		if err != nil {
			return err
		}
		event, err := db.GetEvent(eventID)
		if err != nil {
			return err
		}
		if err := RenderPage(w, config.Data.AdminPath, "editEvent", &Page{Title: "event edition", Event: event}); err != nil {
			return err
		}
	case "del":
		seventID := q.Get("event")
		eventID, err := strconv.Atoi(seventID)
		if err != nil {
			return err
		}
		if err := db.DelEvent(eventID); err != nil {
			return err
		}
		http.Redirect(w, r, "/admin/event", http.StatusSeeOther)
	case "add":
		if r.Method == "POST" {
			layout := "2006-01-02"
			sDate := r.FormValue("date")
			date, err := time.Parse(layout, sDate)
			if err != nil {
				return err
			}
			event := &models.Event{
				EventTitle:    r.FormValue("title"),
				EventOrga:     r.FormValue("organisation"),
				EventLocation: r.FormValue("location"),
				EventDescr:    r.FormValue("description"),
				EventInfo:     r.FormValue("additionalInfo"),
				EventType:     r.FormValue("type"),
				EventDte:      date,
			}
			if err := db.StoreNewEvent(event); err != nil {
				return err
			}
			http.Redirect(w, r, "/admin/event", http.StatusSeeOther)
		}
	case "update":
		if r.Method == "POST" {
			seventID := r.FormValue("id")
			eventID, err := strconv.Atoi(seventID)
			if err != nil {
				return err
			}
			layout := "02-01-2006"
			sDate := r.FormValue("date")
			date, err := time.Parse(layout, sDate)
			if err != nil {
				return err
			}
			event := &models.Event{
				EventTitle:    r.FormValue("title"),
				EventOrga:     r.FormValue("organisation"),
				EventLocation: r.FormValue("location"),
				EventDescr:    r.FormValue("description"),
				EventInfo:     r.FormValue("additionalInfo"),
				EventType:     r.FormValue("type"),
				EventDte:      date,
				EventID:       eventID,
			}
			if err := db.UpdateEvent(event); err != nil {
				return err
			}
			http.Redirect(w, r, "/admin/event", http.StatusSeeOther)
		}
	default:
		if err := RenderPage(w, config.Data.AdminPath, "event", &Page{Title: "event edition", Events: events}); err != nil {
			return err
		}
	}

	return nil
}

func switchCaseInfo(w http.ResponseWriter, r *http.Request) error {
	q := r.URL.Query()
	infos, err := db.GetInfos()
	if err != nil {
		return err
	}
	switch q.Get("do") {
	case "new":
		if err := RenderPage(w, config.Data.AdminPath, "infoAdd", &Page{Title: "info creation"}); err != nil {
			return err
		}
	case "edit":
		sinfoID := q.Get("info")
		infoID, err := strconv.Atoi(sinfoID)
		if err != nil {
			return err
		}
		info, err := db.GetInfo(infoID)
		if err != nil {
			return err
		}
		if err := RenderPage(w, config.Data.AdminPath, "infoEdit", &Page{Title: "info edition", Info: info}); err != nil {
			return err
		}
	case "del":
		sinfoID := q.Get("info")
		infoID, err := strconv.Atoi(sinfoID)
		if err != nil {
			return err
		}
		if err := db.DelInfo(infoID); err != nil {
			return err
		}
		http.Redirect(w, r, "/admin/info", http.StatusSeeOther)
	case "add":
		if r.Method == "POST" {
			layout := "2006-01-02"
			sDate := r.FormValue("date")
			date, err := time.Parse(layout, sDate)
			if err != nil {
				return err
			}
			info := &models.Info{
				InfoTitle:  r.FormValue("title"),
				InfoDescr:  r.FormValue("description"),
				InfoAuthor: r.FormValue("author"),
				InfoDte:    date,
			}
			if err := db.StoreNewInfo(info); err != nil {
				return err
			}
			http.Redirect(w, r, "/admin/info", http.StatusSeeOther)
		}
	case "update":
		if r.Method == "POST" {
			sinfoID := r.FormValue("id")
			infoID, err := strconv.Atoi(sinfoID)
			if err != nil {
				return err
			}
			layout := "02-01-2006"
			sDate := r.FormValue("date")
			date, err := time.Parse(layout, sDate)
			if err != nil {
				return err
			}
			info := &models.Info{
				InfoTitle:  r.FormValue("title"),
				InfoDescr:  r.FormValue("description"),
				InfoAuthor: r.FormValue("author"),
				InfoDte:    date,
				InfoID:     infoID,
			}
			if err := db.UpdateInfo(info); err != nil {
				return err
			}
			http.Redirect(w, r, "/admin/info", http.StatusSeeOther)
		}
	default:
		if err := RenderPage(w, config.Data.AdminPath, "info", &Page{Title: "info edition", Infos: infos}); err != nil {
			return err
		}
	}

	return nil
}
