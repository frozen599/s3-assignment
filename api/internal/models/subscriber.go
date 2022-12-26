package models

import "time"

type Subscriber struct {
	ID           int
	UserID       int
	SubscriberID int
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
}
