package models

import (
	"time"

	"github.com/google/uuid"
)

// Models package for defining DTOs and database models.
type Post struct {
	PostId       uuid.UUID `db:"post_id" json:"post_id"`
	Content      string    `db:"content" json:"content"`
	IncidentDate time.Time `db:"incident_date" json:"incident_date"`
	PostedDate   time.Time `db:"posted_date" json:"posted_date"`
	Latitude     float64   `db:"latitude" json:"latitude"`
	Longtitude   float64   `db:"longitude" json:"longitude"`
	AddressText  string    `db:"address_text" json:"address_text"`
	Location     []byte    `db:"location" json:"location"`
	IsActive     bool      `db:"is_active" json:"is_active"`
}

type Posts struct {
	PostId      uuid.UUID `db:"post_id" json:"post_id"`
	PostedDate  time.Time `db:"posted_date" json:"posted_date"`
	Latitude    float64   `db:"latitude" json:"latitude"`
	Longtitude  float64   `db:"longtitude" json:"longtitude"`
	AddressText string    `db:"address_text" json:"address_text"`
	Location    []byte    `db:"location" json:"location"`
}
