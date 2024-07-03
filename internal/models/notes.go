package models

import (
	"database/sql"
	"time"
)

type Note struct {
	ID        int
	Provider  string
	Patient   string
	Service   string
	StartTime time.Time
	EndTime   time.Time
	Status    string
}

// The reason we have this is so we can use the db connection, without
// having to re-establish it. It is passed in from the main func
type NoteModel struct {
	DB *sql.DB
}

func (*NoteModel) Get(id int) (Note, error) {
	//DB.query to get the note based on the id

	note := Note{
		ID:       000027,
		Provider: "Joe Smith",
		Patient:  "Nina Adams",
		Service:  "Basic",
		Status:   "Pending",
	}

	return note, nil
}
