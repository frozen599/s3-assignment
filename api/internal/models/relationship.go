package models

import "time"

const (
	RelationshipTypeFriend     = "friend"
	RelationshipTypeSubscriber = "subscriber"
	RelationshipTypeBlocking   = "blocking"
)

type Relationship struct {
	ID               int
	UserID1          int `pg:"user_id_1"`
	UserID2          int `pg:"user_id_2"`
	RelationshipType string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        time.Time
}
