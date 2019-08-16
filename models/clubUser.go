package models

import "time"

type ClubUser struct {
	UserID       int
	Name         string
	Firstname    string
	Email        string
	Hash         string
	Role         string
	Rank         string
	RankObtained int
	Licence      string
	MedCert      bool
	EntryDate    int
	CreatedOn    time.Time
}
