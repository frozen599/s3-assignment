package repo

import (
	"github.com/frozen599/s3-assignment/api/internal/config"
	"github.com/frozen599/s3-assignment/api/internal/models"
	"github.com/go-pg/pg/v10"
)

type IRelationshipRepo interface {
	CreateRelationship(friendRelationship models.Relationship) error
	GetFriendList(userID int) ([]models.Relationship, error)
	CheckIfIsBlockingTarget(userId, targetUserId int) (bool, error)
	CheckIfFriendConnectionExists(userID, targetUserID int) (bool, error)
	CanReceiveUpdate(userID int) ([]models.Relationship, error)
}

type relationshipRepo struct {
}

var RelationshipRepo IRelationshipRepo = relationshipRepo{}

func (relationshipRepo) CreateRelationship(friendRelationship models.Relationship) error {
	_, err := config.GetDB().
		Model(&friendRelationship).
		Insert()

	if err != nil {
		return err
	}
	return nil
}

func (relationshipRepo) GetFriendList(userID int) ([]models.Relationship, error) {
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

func (relationshipRepo) CheckIfIsBlockingTarget(userID, targetUserID int) (bool, error) {
	var rela models.Relationship
	err := config.GetDB().
		Model(&rela).
		Where("user_id_1 = ? AND user_id_2 = ?", userID, targetUserID).
		Limit(1).
		Select()

	if err != nil {
		return false, err
	}
	return rela.ID != 0, err
}

func (relationshipRepo) CheckIfFriendConnectionExists(userID, targetUserID int) (bool, error) {
	var rela models.Relationship
	err := config.GetDB().
		Model(&rela).
		Where("user_id_1 = ? AND user_id_2 = ?", userID, targetUserID).
		Limit(1).
		Select()

	if err == pg.ErrNoRows {
		return false, nil
	}
	return rela.ID != 0, err
}

func (relationshipRepo) CanReceiveUpdate(userID int) ([]models.Relationship, error) {
	var ret []models.Relationship
	err := config.GetDB().
		Model(&ret).
		Where("user_id_1 = ? AND relationship_type NOT IN (?)",
			userID,
			pg.In([]string{models.RelationshipTypeBlocking})).
		Select()

	if err == pg.ErrNoRows {
		return nil, nil
	}
	return ret, err
}
