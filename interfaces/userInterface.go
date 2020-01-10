package interfaces

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/riphidon/clubmanager/services"
)

func Edit(w http.ResponseWriter, r *http.Request, member int) error {
	r.ParseForm()
	name := r.FormValue("name")
	firstname := r.FormValue("firstname")
	rank := r.FormValue("belt")
	fmt.Printf("IN EDIT rank: %v\n", rank)
	srankObtained := r.FormValue("rankObtained")
	rankObtained, err := strconv.Atoi(srankObtained)
	if err != nil {
		return err
	}
	fmt.Printf("IN EDIT rankObt: %v\n", rankObtained)
	licence := r.FormValue("licence")
	smedCert := r.FormValue("medCert")
	medCert, err := strconv.ParseBool(smedCert)
	if err != nil {
		return err
	}
	fmt.Printf("IN EDIT certif: %v\n", medCert)
	sentryDate := r.FormValue("entryDate")
	entryDate, err := strconv.Atoi(sentryDate)
	if err != nil {
		return err
	}
	fmt.Printf("IN EDIT entryDate: %v\n", entryDate)
	n := services.UserEdit(name, firstname, rank, licence, medCert, rankObtained, entryDate, member)
	if err := services.UpdateUserProfile(n); err != nil {
		return err
	}
	return nil

}

func AdminEdit(w http.ResponseWriter, r *http.Request, user int) error {
	r.ParseForm()
	rank := r.FormValue("belt")
	sRankObtained := r.FormValue("rankObtained")
	rankObtained, err := strconv.Atoi(sRankObtained)
	if err != nil {
		return err
	}
	licence := r.FormValue("licence")
	group := r.FormValue("group")
	fmt.Printf("group in adminedit: %v\n", group)
	sMedCert := r.FormValue("medCert")
	medCert, err := strconv.ParseBool(sMedCert)
	if err != nil {
		return err
	}
	n := services.AdminUserEdit(rank, licence, group, medCert, rankObtained, user)
	if err := services.AdminUpdate(n); err != nil {
		return err
	}
	return nil
}
