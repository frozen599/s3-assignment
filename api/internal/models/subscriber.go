package models

import "time"

type Subscriber struct {
	ID           int
	UserID       int
	SubscriberId int
	CreatedAt    time.Time
	CreatedBy    int
	UpdatedAt    time.Time
	UpdatedBy    int
	DeletedAt    time.Time
	DeletedBy    int
}
