package server

import (
	"net/http"

	"github.com/riphidon/clubmanager/config"
	"github.com/riphidon/clubmanager/interfaces"
	"github.com/riphidon/clubmanager/utils"
)

//Admin SECTION
func AdminSection(w http.ResponseWriter, r *http.Request) error {
	id := utils.CatchURLData(r, "q")
	userData := interfaces.ViewUserData(r, id, "admin")
	adminPage := utils.SplitPath(r.URL.Path, 1)
	switch adminPage {
	case "profiles":
		id := utils.CatchURLData(r, "q")
		userData := interfaces.ViewUserData(r, id, "admin")
		err := RenderPage(w, config.Data.AdminPath, "profiles", &Page{Title: "profiles", User: userData})
		if err != nil {
			return err
		}
	case "tasks":
		id := utils.CatchURLData(r, "q")
		userData := interfaces.ViewUserData(r, id, "admin")
		err := RenderPage(w, config.Data.AdminPath, "tasks", &Page{Title: "tasks", User: userData})
		if err != nil {
			return err
		}
	case "competitions":
		id := utils.CatchURLData(r, "q")
		userData := interfaces.ViewUserData(r, id, "admin")
		err := RenderPage(w, config.Data.AdminPath, "competitions", &Page{Title: "competitions", User: userData})
		if err != nil {
			return err
		}
	default:
		err := RenderPage(w, config.Data.AdminPath, "admin", &Page{Title: "admin", User: userData})
		if err != nil {
			return err
		}
	}

	return nil
}
