package repository

import (
	"github.com/frozen599/s3-assignment/api/internal/config"
	"github.com/frozen599/s3-assignment/api/internal/models"
	"github.com/go-pg/pg/v10"
)

func CreateFriendConnection(friendRelationship models.Relationship) error {
	_, err := config.GetDB().
		Model(&friendRelationship).
		Insert()

	if err != nil {
		return err
	}
	return nil
}

func GetFriendList(userID int) ([]models.Relationship, error) {
	var ret []models.Relationship
	err := config.GetDB().
		Model(&ret).
		Where("user_id_1 = ? AND relationship_type = ?", userID, models.RelationshipTypeFriend).
		Select()
	if err == pg.ErrNoRows {
		return nil, nil
	}
	return ret, err
}
