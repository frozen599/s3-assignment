package usescase

import (
	"time"

	"github.com/frozen599/s3-assignment/api/internal/models"
	"github.com/frozen599/s3-assignment/api/internal/pkg"
	"github.com/frozen599/s3-assignment/api/internal/repository"
)

type FriendUseCase interface {
	CreateFriendConnection(firsrUserEmail, secondUserEmail string) error
	GetFriendList(email string) ([]models.Relationship, error)
	GetMutualFriendList(firsrUserEmail, secondUserEmail string) ([]models.Relationship, error)
}

type friendUseCase struct {
}

func NewFriendUseCase() FriendUseCase {
	return friendUseCase{}
}

func (f friendUseCase) CreateFriendConnection(firstUserEmail, secondUserEmail string) error {
	firstUser, err := repository.GetUserByEmail(firstUserEmail)
	if err != nil {
		return err
	}
	secondUser, err := repository.GetUserByEmail(secondUserEmail)
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
	err = repository.CreateRelationship(friendRelationShip)
	if err != nil {
		return err
	}
	return nil
}

func (f friendUseCase) GetFriendList(email string) ([]models.Relationship, error) {
	user, err := repository.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	friendList, err := repository.GetFriendList(user.ID)
	if err != nil {
		return nil, err
	}
	return friendList, err
}

func (f friendUseCase) GetMutualFriendList(firstUserEmail, secondUserEmail string) ([]models.Relationship, error) {
	firstUser, err := repository.GetUserByEmail(firstUserEmail)
	if err != nil {
		return nil, err
	}
	secondUser, err := repository.GetUserByEmail(secondUserEmail)
	if err != nil {
		return nil, err
	}

	firstUserFriendList, err := repository.GetFriendList(firstUser.ID)
	if err != nil {
		return nil, err
	}
	secondUserFriendList, err := repository.GetFriendList(secondUser.ID)
	if err != nil {
		return nil, err
	}

	mutualFriendList := pkg.GetMutualFriendList(firstUserFriendList, secondUserFriendList)
	return mutualFriendList, err
}