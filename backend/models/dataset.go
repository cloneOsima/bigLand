package models

// Models package for defining DTOs and database models.

type Post struct {
	Id                 int    `db:"id" json:"id"`
	AccidentDate       string `db:"accident_date" json:"accident_date"`
	AccidentCause      string `db:"accident_cause" json:"accident_cause"`
	CreatedAt          string `db:"created_at" json:"created_at"`
	RegistratedAddress string `db:"registed_address" json:"registed_address"`
}

type EntirePost struct {
	Id                 int    `db:"id" json:"id"`
	CreatedAt          string `db:"created_at" json:"created_at"`
	RegistratedAddress string `db:"registed_address" json:"registed_address"`
}
