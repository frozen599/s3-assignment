package models

import "time"

type Relationship struct {
	ID               int
	UserID1          int
	UserID2          int
	RelationshipType string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        time.Time
}
