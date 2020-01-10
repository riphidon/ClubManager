package models

import "time"

type Info struct {
	InfoID     int
	InfoTitle  string
	InfoDescr  string
	InfoAuthor string
	InfoDte    time.Time
}
