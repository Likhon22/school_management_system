package models

import (
	"time"
)

type Role string

const (
	RoleAdmin   Role = "admin"
	RoleTeacher Role = "teacher"
	RoleStudent Role = "student"
	RoleExec    Role = "exec"
)

type Exec struct {
	ID                int       `json:"id" db:"id"`
	FirstName         string    `json:"first_name,omitempty" db:"first_name"`
	LastName          string    `json:"last_name,omitempty" db:"last_name"`
	Email             string    `json:"email,omitempty" db:"email"`
	Username          string    `json:"username,omitempty" db:"username"`
	Password          string    `json:"password,omitempty" db:"password"`
	PasswordChangedAt time.Time `json:"password_changed_at,omitempty" db:"password_changed_at"`
	PasswordResetCode string    `json:"password_reset_code,omitempty" db:"password_reset_code"`
	Role              Role      `json:"role,omitempty" db:"role"`
	CreatedAt         time.Time `json:"created_at" db:"created_at"`
	UpdatedAt         time.Time `json:"updated_at" db:"updated_at"`
}
