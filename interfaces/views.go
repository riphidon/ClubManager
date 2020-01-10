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
		fmt.Printf("error occured: %v\n", err)
	}
	return belts
}

func ViewWeightList() []string {
	weights, err := db.WeightList()
	if err != nil {
		fmt.Printf("error occured: %v\n", err)
	}
	return weights
}

func ViewAgeList() []string {
	ages, err := db.AgeList()
	if err != nil {
		fmt.Printf("error occured: %v\n", err)
	}
	return ages
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

func ListAllUsers() []*models.ClubUser {
	userList, err := db.AllUsers()
	if err != nil {
		fmt.Printf("error occured in ListAllUser: %v\n", err)
	}
	return userList
}

func FindUserById(id int) models.ClubUser {
	// ID, err := strconv.P(id)
	// if err != nil {
	// 	fmt.Printf("error occured parsing data: %v\n", err)
	// }
	user, err := db.UsersById(id)
	if err != nil {
		fmt.Printf("error occured in FindUserById: %v\n", err)
	}
	return user
}

func FindUserByRank(belt string) []*models.ClubUser {
	users, err := db.UsersByRank(belt)
	if err != nil {
		fmt.Printf("error occured in FindUserByRank: %v\n", err)
	}
	return users

}
