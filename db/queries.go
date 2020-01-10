package db

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/riphidon/clubmanager/models"
)

//Create
func StoreNewUser(u models.ClubUser) error {
	query := `INSERT INTO club_user (name, firstname, email, hash, role,
			 rank, rank_obtained, med_cert ,entry_date, licence)
			 VALUES ($1, $2, $3, $4 ,$5 , $6, $7, $8, $9, $10)`
	_, err := DB.Exec(query, u.Name, u.Firstname, u.Email, u.Hash, u.Role, u.Rank, u.RankObtained, u.MedCert, u.EntryDate, u.Licence)
	if err != nil {
		fmt.Println("registering error")
		return errors.Wrap(err, "Can't execute sql statement")
	}
	fmt.Println("registering")
	return nil
}

//Read
func BeltList() ([]string, error) {
	var beltList []string
	query := `SELECT color
			FROM belts`
	rows, err := DB.Query(query)
	if err != nil {
		return beltList, errors.Wrap(err, "can't query Database")
	}
	for rows.Next() {
		var beltColor string
		if err := rows.Scan(&beltColor); err != nil {
			return beltList, errors.Wrap(err, "can't perform scan")
		}
		beltList = append(beltList, beltColor)
	}
	return beltList, nil
}

func WeightList() ([]string, error) {
	var weightList []string
	query := `SELECT name
			FROM weight_cat`
	rows, err := DB.Query(query)
	if err != nil {
		return weightList, errors.Wrap(err, "can't query Database")
	}
	for rows.Next() {
		var weightCat string
		if err := rows.Scan(&weightCat); err != nil {
			return weightList, errors.Wrap(err, "can't perform scan")
		}
		weightList = append(weightList, weightCat)
	}
	return weightList, nil
}

func AgeList() ([]string, error) {
	var ageList []string
	query := `SELECT name
			FROM age_cat`
	rows, err := DB.Query(query)
	if err != nil {
		return ageList, errors.Wrap(err, "can't query Database")
	}
	for rows.Next() {
		var ageCat string
		if err := rows.Scan(&ageCat); err != nil {
			return ageList, errors.Wrap(err, "can't perform scan")
		}
		ageList = append(ageList, ageCat)
	}
	return ageList, nil
}

func UserData(id int) (*models.ClubUser, error) {
	var data = new(models.ClubUser)
	query := `SELECT user_id,
					 name,
  					firstname,
					  rank,
					  rank_obtained,
  					licence,
					  med_cert,
					  entry_date
  			FROM club_user
  			WHERE user_id = $1`
	if err := DB.QueryRow(query, id).Scan(&data.UserID, &data.Name,
		&data.Firstname, &data.Rank, &data.RankObtained, &data.Licence,
		&data.MedCert, &data.EntryDate); err != nil {
		return data, errors.Wrap(err, "Couldn't execute sql statement")
	}
	return data, nil
}

//Update
func ProfileByUser(u models.ClubUser) error {
	query := `UPDATE club_user
			SET name = $1, firstname = $2, rank = $3, rank_obtained = $4,
			 licence = $5, med_cert = $6, entry_date = $7
			 WHERE user_id = $8`
	_, err := DB.Exec(query, u.Name, u.Firstname, u.Rank, u.RankObtained, u.Licence, u.MedCert, u.EntryDate, u.UserID)
	if err != nil {
		return errors.Wrap(err, "Can't execute sql statement")
	}
	return nil
}

//Delete

//---SEARCH---
func UserExists(email string) (bool, error) {
	var exists bool
	query := ` SELECT EXISTS 
			(SELECT 1 FROM club_user 
			WHERE email = $1);`
	err := DB.QueryRow(query, email).Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "Couldn't execute sql statement")
	}
	fmt.Printf("userExists: %v\n", exists)
	return exists, nil
}

func AllUsers() ([]*models.ClubUser, error) {
	query := `
		SELECT user_id, name, firstname, rank, licence, med_cert
		FROM club_user`
	rows, err := DB.Query(query)
	if err != nil {
		return nil, errors.Wrap(err, "can't query Database")
	}
	defer rows.Close()
	userList := make([]*models.ClubUser, 0)
	for rows.Next() {
		user := new(models.ClubUser)
		err := rows.Scan(&user.UserID, &user.Name, &user.Firstname, &user.Rank, &user.Licence, &user.MedCert)
		if err != nil {
			return nil, errors.Wrap(err, "can't perform scan")
		}
		userList = append(userList, user)
	}
	err = rows.Err()
	if err != nil {
		return nil, errors.Wrap(err, "Additional errors while scanning database rows")
	}
	return userList, nil
}

func UsersById(id int) (models.ClubUser, error) {
	var user models.ClubUser
	query := `
		SELECT user_id, name, firstname, rank, licence, med_cert, role, entry_date, rank_obtained
		FROM club_user
		WHERE user_id = $1`
	err := DB.QueryRow(query, id).Scan(&user.UserID, &user.Name, &user.Firstname, &user.Rank, &user.Licence, &user.MedCert, &user.Role, &user.EntryDate, &user.RankObtained)
	if err != nil {
		return user, errors.Wrap(err, "Couldn't execute sql statement")
	}
	return user, nil
}

func Hash(email string) (string, error) {
	hash := ""
	query := `SELECT hash
			FROM club_user 
			WHERE email = $1`
	if err := DB.QueryRow(query, email).Scan(&hash); err != nil {
		return hash, errors.Wrap(err, "Couldn't execute sql statement")
	}
	return hash, nil
}

func ID(email string) (string, error) {
	var ID string
	query := `SELECT user_id
			FROM club_user 
			WHERE email = $1`
	if err := DB.QueryRow(query, email).Scan(&ID); err != nil {
		return ID, errors.Wrap(err, "Couldn't execute sql statement")
	}
	return ID, nil
}
