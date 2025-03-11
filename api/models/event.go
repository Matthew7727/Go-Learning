package models

import "time"

type Event struct {

	ID int
	Name string `binding:"required"`
	Desctiption string `binding:"required"`
	Location string `binding:"required"`
	DateTime time.Time `binding:"required"`
	UserID int
}

var events = []Event{}

func (e Event) Save() {
	//TODO : add it to database
	events = append(events, e)
}

func GetEvents() []Event {
	return events
}