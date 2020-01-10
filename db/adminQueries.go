package db

import (
	"github.com/pkg/errors"
	"github.com/riphidon/clubmanager/models"
)

//Create

func StoreNewTask(t models.Task) error {
	query := `INSERT INTO task (content, author)
			VALUES ($1, $2)`
	_, err := DB.Exec(query, t.Content, t.Author)
	if err != nil {
		return errors.Wrap(err, "Can't execute sql statement")
	}
	return nil
}

func StoreNewEvent(e *models.Event) error {
	query := `INSERT INTO event (title, orga, location, descr, infos, type, dte)
			VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := DB.Exec(query, e.EventTitle, e.EventOrga, e.EventLocation, e.EventDescr, e.EventInfo, e.EventType, e.EventDte)
	if err != nil {
		return errors.Wrap(err, "Can't execute sql statement")
	}
	return nil
}

func StoreNewInfo(i *models.Info) error {
	query := `INSERT INTO info (title, descr, author, dte)
			VALUES ($1, $2, $3, $4)`
	_, err := DB.Exec(query, i.InfoTitle, i.InfoDescr, i.InfoAuthor, i.InfoDte)
	if err != nil {
		return errors.Wrap(err, "Can't execute sql statement")
	}
	return nil
}

//Read
func AdminData(id int) (*models.ClubUser, error) {
	var data = new(models.ClubUser)
	query := `SELECT user_id,
					 name,
  					firstname,
					  rank,
  					licence
  			FROM club_user
  			WHERE user_id = $1`
	if err := DB.QueryRow(query, id).Scan(&data.UserID, &data.Name,
		&data.Firstname, &data.Rank, &data.Licence); err != nil {
		return data, errors.Wrap(err, "Couldn't execute sql statement")
	}
	return data, nil
}

func ListByParam(param, data string) ([]*models.ClubUser, error) {
	var query string
	switch param {
	case "name":
		query = `
		SELECT user_id, name, firstname, rank
		FROM club_user
		WHERE name = $1`
	case "rank":
		query = `
		SELECT user_id, name, firstname, rank
		FROM club_user
		WHERE rank = $1`
	}
	rows, err := DB.Query(query, data)
	if err != nil {
		return nil, errors.Wrap(err, "can't query Database")
	}
	defer rows.Close()
	userList := make([]*models.ClubUser, 0)
	for rows.Next() {
		user := new(models.ClubUser)
		err := rows.Scan(&user.UserID, &user.Name, &user.Firstname, &user.Rank)
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

func GetEvents() ([]*models.Event, error) {
	query := `SELECT * FROM event`
	rows, err := DB.Query(query)
	if err != nil {
		return nil, errors.Wrap(err, "Can't query database")
	}
	defer rows.Close()
	events := make([]*models.Event, 0)
	for rows.Next() {
		event := new(models.Event)
		err := rows.Scan(&event.EventID, &event.EventTitle, &event.EventOrga, &event.EventLocation, &event.EventDescr, &event.EventInfo, &event.EventType, &event.EventDte)
		if err != nil {
			return nil, errors.Wrap(err, "can't scan database rows")
		}
		events = append(events, event)
	}
	err = rows.Err()
	if err != nil {
		return nil, errors.Wrap(err, "Additional errors while scanning database rows")
	}
	return events, nil
}

func GetEvent(eventID int) (*models.Event, error) {
	event := new(models.Event)
	query := `SELECT * FROM event
				WHERE event_id = $1`
	if err := DB.QueryRow(query, eventID).Scan(&event.EventID, &event.EventTitle, &event.EventOrga, &event.EventLocation, &event.EventDescr, &event.EventInfo, &event.EventType, &event.EventDte); err != nil {
		return event, errors.Wrap(err, "Couldn't execute sql statement")
	}
	return event, nil
}

func GetInfos() ([]*models.Info, error) {
	query := `SELECT * FROM info`
	rows, err := DB.Query(query)
	if err != nil {
		return nil, errors.Wrap(err, "Can't query database")
	}
	defer rows.Close()
	infos := make([]*models.Info, 0)
	for rows.Next() {
		info := new(models.Info)
		err := rows.Scan(&info.InfoID, &info.InfoTitle, &info.InfoDescr, &info.InfoAuthor, &info.InfoDte)
		if err != nil {
			return nil, errors.Wrap(err, "can't scan database rows")
		}
		infos = append(infos, info)
	}
	err = rows.Err()
	if err != nil {
		return nil, errors.Wrap(err, "Additional errors while scanning database rows")
	}
	return infos, nil
}

func GetInfo(infoID int) (*models.Info, error) {
	info := new(models.Info)
	query := `SELECT * FROM info
				WHERE info_id = $1`
	if err := DB.QueryRow(query, infoID).Scan(&info.InfoID, &info.InfoTitle, &info.InfoDescr, &info.InfoAuthor, &info.InfoDte); err != nil {
		return info, errors.Wrap(err, "Couldn't execute sql statement")
	}
	return info, nil
}

//Update
func ProfileByAdmin(u models.ClubUser) error {
	query := `UPDATE club_user
			SET role = $1, rank = $2, rank_obtained = $3,
			 licence = $4, med_cert = $5
			 WHERE user_id = $6`
	_, err := DB.Exec(query, u.Role, u.Rank, u.RankObtained, u.Licence, u.MedCert, u.UserID)
	if err != nil {
		return errors.Wrap(err, "Can't execute sql statement")
	}
	return nil
}

func CompByAdmin(c models.Competition) error {
	query := `UPDATE competitionu
			SET name = $1, orga = $2, location = $3, dte = $4
			WHERE comp_id = $5`
	_, err := DB.Exec(query, c.Name, c.Orga, c.Location, c.Dte, c.CompID)
	if err != nil {
		return errors.Wrap(err, "Can't execute sql statement")
	}
	return nil
}

func TaskByAdmin(t models.Task) error {
	query := `UPDATE task
			SET content = $1, author = $2
			WHERE task_id = $5`
	_, err := DB.Exec(query, t.Content, t.Author, t.TaskID)
	if err != nil {
		return errors.Wrap(err, "Can't execute sql statement")
	}
	return nil
}

func UpdateEvent(e *models.Event) error {
	query := `UPDATE event
				SET title = $1,
					orga = $2,
					location = $3,
					descr = $4,
					infos = $5,
					type = $6,
					dte = $7
				WHERE event_id = $8`
	_, err := DB.Exec(query, e.EventTitle, e.EventOrga, e.EventLocation, e.EventDescr, e.EventInfo, e.EventType, e.EventDte, e.EventID)
	if err != nil {
		return errors.Wrap(err, "Can't execute sql statement")
	}
	return nil
}

func UpdateInfo(i *models.Info) error {
	query := `UPDATE info
				SET title = $1,
					descr = $2,
					author = $3,
					dte = $4
				WHERE info_id = $5`
	_, err := DB.Exec(query, i.InfoTitle, i.InfoDescr, i.InfoAuthor, i.InfoDte, i.InfoID)
	if err != nil {
		return errors.Wrap(err, "Can't execute sql statement")
	}
	return nil
}

//Delete
func DelEvent(id int) error {
	query := `DELETE FROM event
				WHERE event_id = $1`
	_, err := DB.Exec(query, id)
	if err != nil {
		return errors.Wrap(err, "Can't execute sql statement")
	}
	return nil
}

func DelInfo(id int) error {
	query := `DELETE FROM info
				WHERE info_id = $1`
	_, err := DB.Exec(query, id)
	if err != nil {
		return errors.Wrap(err, "Can't execute sql statement")
	}
	return nil
}

//---SEARCH---

func UsersByRank(belt string) ([]*models.ClubUser, error) {
	query := `
		SELECT user_id, name, firstname, rank, licence, med_cert, role
		FROM club_user
		WHERE rank = $1`
	rows, err := DB.Query(query, belt)
	if err != nil {
		return nil, errors.Wrap(err, "can't query Database")
	}
	defer rows.Close()
	userList := make([]*models.ClubUser, 0)
	for rows.Next() {
		user := new(models.ClubUser)
		err := rows.Scan(&user.UserID, &user.Name, &user.Firstname, &user.Rank, &user.Licence, &user.MedCert, &user.Role)
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

func Role(email string) (string, error) {
	group := ""
	query := ` SELECT role
 			FROM club_user
 			WHERE email = $1`
	if err := DB.QueryRow(query, email).Scan(&group); err != nil {
		return group, errors.Wrap(err, "Couldn't execute sql statement")
	}
	return group, nil
}
