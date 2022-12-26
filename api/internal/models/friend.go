package models

import "time"

type Friend struct {
	ID        int
	UserID1   int
	UserID2   int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
