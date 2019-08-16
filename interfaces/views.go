package interfaces

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/riphidon/clubmanager/db"
	"github.com/riphidon/clubmanager/models"
)

func ViewBeltList() []string {
	belts, err := db.BeltList()
	if err != nil {
		fmt.Printf("error occured: %v", err)
	}
	return belts
}

func ViewUserData(r *http.Request, data string, group string) models.ClubUser {
	id := data
	ID, _ := strconv.Atoi(id)
	fmt.Printf("id: %v", ID)
	if group == "user" {
		userData, err := db.UserData(ID)
		if err != nil {
			fmt.Printf("error occured in ViewUserData: %v\n", err)
		}
		return *userData
	}
	userData, err := db.AdminData(ID)
	if err != nil {
		fmt.Printf("error occured in AdminData: %v\n", err)
	}
	return *userData

}