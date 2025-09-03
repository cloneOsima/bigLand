package models

type Post struct {
	Id                 int    `db:"id" json:"id"`
	AccidentDate       string `db:"accident_date" json:"accident_date"`
	AccidentCause      string `db:"accident_cause" json:"accident_cause"`
	RegistrationDate   string `db:"registration_date" json:"registration_date"`
	RegistratedAddress string `db:"registed_address" json:"registed_address"`
}
