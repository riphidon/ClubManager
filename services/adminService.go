package services

import (
	"fmt"

	"github.com/riphidon/clubmanager/db"
	"github.com/riphidon/clubmanager/models"
)

func AdminUserEdit(rank, licence, group string, medCert bool, rankObtained, user int) models.ClubUser {
	return models.ClubUser{
		UserID:       user,
		Rank:         rank,
		RankObtained: rankObtained,
		Licence:      licence,
		MedCert:      medCert,
		Role:         group,
	}
}

func AdminUpdate(u models.ClubUser) error {
	if err := db.ProfileByAdmin(u); err != nil {
		fmt.Printf("Error in AdminUpdate: %v", err)
		return err
	}
	return nil
}

func CreateEvent(e *models.Event) error {
	if err := db.StoreNewEvent(e); err != nil {
		return err
	}
	return nil
}

func CreateInfo(i *models.Info) error {
	if err := db.StoreNewInfo(i); err != nil {
		return err
	}
	return nil
}
