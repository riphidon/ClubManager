package models

import "time"

type Competition struct {
	CompID   int
	Name     string
	Orga     string
	Location string
	Dte      time.Time
}
