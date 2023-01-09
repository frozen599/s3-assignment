package repo

import (
	"github.com/frozen599/s3-assignment/api/internal/models"
	"github.com/go-pg/pg/v10"
)

type RelationshipRepo interface {
	CreateRelationship(friendRelationship models.Relationship) error
	GetFriendList(userID int) ([]int, error)
	CheckIfIsBlockingOrBlocked(userID, targetUserID int) (bool, error)
	CheckIfFriendConnectionExists(userID, targetUserID int) (bool, error)
	CheckIfAlreadySubscribing(userID, targetUserID int) (bool, error)
	GetFollowers(userID int) ([]int, error)
}

type relationshipRepo struct {
	db *pg.DB
}

func NewRelationshipRepo(dbInstance *pg.DB) RelationshipRepo {
	return relationshipRepo{db: dbInstance}
}

func (r relationshipRepo) CreateRelationship(friendRelationship models.Relationship) error {
	_, err := r.db.
		Model(&friendRelationship).
		Insert()

	if err != nil {
		return err
	}
	return nil
}

func (r relationshipRepo) GetFriendList(userID int) ([]int, error) {
	var ret []int
	sql :=
		`
	SELECT user_id_2
	FROM relationships
	WHERE (user_id_1 = ? AND relationship_type = 'friend')
	UNION
	SELECT user_id_1
	FROM relationships
	WHERE (user_id_2 = ? AND relationship_type = 'friend')
	`
	_, err := r.db.Query(&ret, sql, userID, userID)
	if err == pg.ErrNoRows {
		return nil, nil
	}
	return ret, err
}

func (r relationshipRepo) CheckIfIsBlockingOrBlocked(userID, targetUserID int) (bool, error) {
	var rela models.Relationship
	err := r.db.
		Model(&rela).
		Where("user_id_1 = ? AND user_id_2 = ? AND relationship_type = ?", userID, targetUserID, models.RelationshipTypeBlocking).
		WhereOr("user_id_2 = ? AND user_id_1 = ? AND relationship_type = ?", userID, targetUserID, models.RelationshipTypeBlocking).
		Limit(1).
		Select()
	if err == pg.ErrNoRows {
		return false, nil
	}
	return rela.ID != 0, err
}

func (r relationshipRepo) CheckIfFriendConnectionExists(userID, targetUserID int) (bool, error) {
	var rela models.Relationship
	err := r.db.
		Model(&rela).
		Where("user_id_1 = ? AND user_id_2 = ? AND relationship_type = ?", userID, targetUserID, models.RelationshipTypeFriend).
		WhereOr("user_id_2 = ? AND user_id_1 = ? AND relationship_type = ?", userID, targetUserID, models.RelationshipTypeFriend).
		Limit(1).
		Select()
	if err == pg.ErrNoRows {
		return false, nil
	}
	return rela.ID != 0, err
}

func (r relationshipRepo) CheckIfAlreadySubscribing(userID, targetUserID int) (bool, error) {
	var rela models.Relationship
	err := r.db.
		Model(&rela).
		Where("user_id_1 = ? AND user_id_2 = ? AND relationship_type = ?", userID, targetUserID, models.RelationshipTypeSubscriber).
		Limit(1).
		Select()
	if err == pg.ErrNoRows {
		return false, nil
	}
	return rela.ID != 0, err
}

func (r relationshipRepo) GetFollowers(userID int) ([]int, error) {
	var ret []int
	sql :=
		`
	SELECT user_id_1
	FROM relationships
	WHERE user_id_2 = ?
		AND relationship_type = ?
	`
	_, err := r.db.Query(&ret, sql, userID, models.RelationshipTypeSubscriber)
	if err == pg.ErrNoRows {
		return nil, nil
	}
	return ret, err
}
