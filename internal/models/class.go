package models

import "time"

type Class struct {
	ID        int       `json:"id,omitempty" db:"id"`
	Name      string    `json:"name,omitempty" db:"name"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
