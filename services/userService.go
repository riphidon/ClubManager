package services

import (
	"fmt"

	"github.com/riphidon/clubmanager/db"
	"github.com/riphidon/clubmanager/models"
)

func UserEdit(name, firstname, rank, licence string, medCert bool, rankObtained, entryDate, user int) models.ClubUser {
	return models.ClubUser{
		UserID:       user,
		Name:         name,
		Firstname:    firstname,
		Rank:         rank,
		RankObtained: rankObtained,
		Licence:      licence,
		MedCert:      medCert,
		EntryDate:    entryDate,
	}
}

func UpdateUserProfile(u models.ClubUser) error {
	if err := db.ProfileByUser(u); err != nil {
		return err
	}
	fmt.Println("UpdateUserProfile no error")
	return nil
}
