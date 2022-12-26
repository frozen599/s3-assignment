package models

import (
	"time"
)

type User struct {
	ID           int
	Email        string
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
}
