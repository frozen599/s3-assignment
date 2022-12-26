package models

import "time"

type Subscriber struct {
	ID              int
	RequestorUserID int
	TargetUserID    int
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       time.Time
}
