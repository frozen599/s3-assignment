package repository

import (
	"github.com/frozen599/s3-assignment/api/internal/config"
	"github.com/frozen599/s3-assignment/api/internal/models"
	"github.com/go-pg/pg/v10"
)

type RelationshipRepository interface {
	CreateRelationship(friendRelationship models.Relationship) error
	GetFriendList(userID int) ([]models.Relationship, error)
	CheckIfBlocking(userId, targetUserId int) (bool, error)
}

type relationshipRepository struct {
}

func NewRelationshipRepository() RelationshipRepository {
	return relationshipRepository{}
}

func (r relationshipRepository) CreateRelationship(friendRelationship models.Relationship) error {
	_, err := config.GetDB().
		Model(&friendRelationship).
		Insert()

	if err != nil {
		return err
	}
	return nil
}

func (r relationshipRepository) GetFriendList(userID int) ([]models.Relationship, error) {
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

func (r relationshipRepository) CheckIfBlocking(userId, targetUserId int) (bool, error) {
	var rela models.Relationship
	err := config.GetDB().
		Model(&rela).
		Where("user_id_1 = ? AND user_id_2 = ?", userId, targetUserId).
		Limit(1).
		Select()

	if err != nil {
		return false, err
	}
	return rela.ID != 0, err
}
