package models

import (
	"database/sql"
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
	ID                        int            `json:"id" db:"id"`
	FirstName                 string         `json:"first_name,omitempty" db:"first_name"`
	LastName                  string         `json:"last_name,omitempty" db:"last_name"`
	Email                     string         `json:"email,omitempty" db:"email"`
	Username                  string         `json:"username,omitempty" db:"username"`
	Password                  string         `json:"password,omitempty" db:"password"`
	PasswordChangedAt         sql.NullTime   `json:"password_changed_at,omitempty" db:"password_changed_at"`
	PasswordResetToken        sql.NullString `json:"password_reset_token,omitempty" db:"password_reset_token"`
	PasswordResetTokenExpires sql.NullString `json:"password_reset_token_expire,omitempty" db:"password_reset_token_expire"`
	Role                      Role           `json:"role,omitempty" db:"role"`
	CreatedAt                 time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt                 time.Time      `json:"updated_at" db:"updated_at"`
}

type ResExec struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}
