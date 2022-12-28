package controller

import (
	"time"

	"github.com/frozen599/s3-assignment/api/internal/models"
	"github.com/frozen599/s3-assignment/api/internal/pkg"
	"github.com/frozen599/s3-assignment/api/internal/repository"
)

type FriendController interface {
	CreateFriendConnection(firsrUserEmail, secondUserEmail string) error
	GetFriendList(email string) ([]models.User, error)
	GetMutualFriendList(firsrUserEmail, secondUserEmail string) ([]models.User, error)
}

type friendController struct {
}

func NewFriendController() FriendController {
	return friendController{}
}

func (f friendController) CreateFriendConnection(firstUserEmail, secondUserEmail string) error {
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

func (f friendController) GetFriendList(email string) ([]models.User, error) {
	user, err := repository.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	friendListRelationships, err := repository.GetFriendList(user.ID)
	if err != nil {
		return nil, err
	}

	var friendIDs []int
	for _, friend := range friendListRelationships {
		friendIDs = append(friendIDs, friend.ID)
	}
	friends, err := repository.GetUserByIds(friendIDs)
	if err != nil {
		return nil, err
	}
	return friends, nil
}

func (f friendController) GetMutualFriendList(firstUserEmail, secondUserEmail string) ([]models.User, error) {
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

	var friendIDs []int
	for _, friend := range mutualFriendList {
		friendIDs = append(friendIDs, friend.ID)
	}
	friends, err := repository.GetUserByIds(friendIDs)
	if err != nil {
		return nil, err
	}
	return friends, nil
}
