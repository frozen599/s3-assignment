package models

import "time"

type Friend struct {
	ID        int       `db:"id"`
	UserID1   int       `db:"user_id_1"`
	UserID2   int       `db:"user_id_2"`
	CreatedAt time.Time `db:"created_at"`
	DeletedAt time.Time `db:"deleted_at"`
}
