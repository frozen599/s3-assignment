package repository

import (
	"fmt"
	"time"

	"github.com/frozen599/s3-assignment/api/internal/config"
	"github.com/frozen599/s3-assignment/api/internal/models"
	"github.com/go-pg/pg/v10"
)

func CreateFriend(firstUserEmail, secondUserEmail string) error {
	firstUser, err := GetUserByEmail(firstUserEmail)
	if err != nil {
		return err
	}
	secondUser, err := GetUserByEmail(secondUserEmail)
	if err != nil {
		return err
	}

	friendRelationShip := models.Relationship{
		UserID1:          firstUser.ID,
		UserID2:          secondUser.ID,
		RelationshipType: models.RelationshipTypeFriend,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	_, err = config.GetDB().
		Model(&friendRelationShip).
		Insert()

	if err != nil {
		return err
	}
	return nil
}

func GetFriendList(email string) ([]models.Relationship, error) {
	user, err := GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	var ret []models.Relationship
	err = config.GetDB().
		Model(&ret).
		Where("user_id_1 = ? AND relationship_type = ?", user.ID, models.RelationshipTypeFriend).
		Select()

	if err == pg.ErrNoRows {
		fmt.Println("ERRR")
		return nil, nil
	}
	return ret, err
}
