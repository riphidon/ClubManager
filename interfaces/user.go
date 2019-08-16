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
