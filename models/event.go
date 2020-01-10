package models

import "time"

type Event struct {
	EventID       int
	EventTitle    string
	EventOrga     string
	EventLocation string
	EventDescr    string
	EventInfo     string
	EventType     string
	EventDte      time.Time
}
