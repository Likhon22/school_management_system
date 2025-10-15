package models

import "time"

type Teacher struct {
	ID        int       `json:"id" db:"id"`
	Email     string    `json:"email" db:"email"`
	FirstName string    `json:"first_name,omitempty" db:"first_name"`
	LastName  string    `json:"last_name,omitempty" db:"last_name"`
	ClassID   int       `json:"class_id,omitempty" db:"class_id"` // <- change here
	Subject   string    `json:"subject,omitempty" db:"subject"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
