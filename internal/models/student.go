package models

import "time"

type Student struct {
	ID        int       `json:"id,omitempty" db:"id"`
	Email     string    `json:"email,omitempty" db:"email"`
	FirstName string    `json:"first_name,omitempty" db:"first_name"`
	LastName  string    `json:"last_name,omitempty" db:"last_name"`
	ClassID   int       `json:"class_id,omitempty" db:"class_id"` // <- change here
	CreatedAt time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" db:"updated_at"`
}
