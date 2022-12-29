package controller

import (
	"time"

	"github.com/frozen599/s3-assignment/api/internal/models"
	"github.com/frozen599/s3-assignment/api/internal/repository"
	"github.com/frozen599/s3-assignment/api/internal/utils"
)

type FriendController interface {
	CreateFriendConnection(firsrUserEmail, secondUserEmail string) error
	GetFriendList(email string) ([]models.User, error)
	GetMutualFriendList(firsrUserEmail, secondUserEmail string) ([]models.User, error)
}

type friendController struct {
	userRepo repository.UserRepository
	relaRepo repository.RelationshipRepository
}

func NewFriendController() FriendController {
	return friendController{
		userRepo: repository.NewUserRepository(),
		relaRepo: repository.NewRelationshipRepository(),
	}
}

func (fc friendController) CreateFriendConnection(firstUserEmail, secondUserEmail string) error {
	firstUser, err := fc.userRepo.GetUserByEmail(firstUserEmail)
	if err != nil {
		return err
	}
	secondUser, err := fc.userRepo.GetUserByEmail(secondUserEmail)
	if err != nil {
		return err
	}

	existedFriendship, err := fc.relaRepo.CheckIfFriendConnectionExists(firstUser.ID, secondUser.ID)
	if err != nil {
		return err
	}
	if existedFriendship {
		return utils.ErrFriendshipAlreadyExists
	}

	isBlockingTarget, err := fc.relaRepo.CheckIfIsBlockingTarget(firstUser.ID, secondUser.ID)
	if err != nil {
		return err
	}
	if isBlockingTarget {
		return utils.ErrCurrentUserIsBlockingTarget
	}

	friendRelationShip := models.Relationship{
		UserID1:          firstUser.ID,
		UserID2:          secondUser.ID,
		RelationshipType: models.RelationshipTypeFriend,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}
	err = fc.relaRepo.CreateRelationship(friendRelationShip)
	if err != nil {
		return err
	}
	return nil
}

func (fc friendController) GetFriendList(email string) ([]models.User, error) {
	user, err := fc.userRepo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	friendListRelationships, err := fc.relaRepo.GetFriendList(user.ID)
	if err != nil {
		return nil, err
	}

	var friendIDs []int
	for _, friend := range friendListRelationships {
		friendIDs = append(friendIDs, friend.ID)
	}
	friends, err := fc.userRepo.GetUserByIDs(friendIDs)
	if err != nil {
		return nil, err
	}
	return friends, nil
}

func (fc friendController) GetMutualFriendList(firstUserEmail, secondUserEmail string) ([]models.User, error) {
	firstUser, err := fc.userRepo.GetUserByEmail(firstUserEmail)
	if err != nil {
		return nil, err
	}
	secondUser, err := fc.userRepo.GetUserByEmail(secondUserEmail)
	if err != nil {
		return nil, err
	}

	firstUserFriendList, err := fc.relaRepo.GetFriendList(firstUser.ID)
	if err != nil {
		return nil, err
	}
	secondUserFriendList, err := fc.relaRepo.GetFriendList(secondUser.ID)
	if err != nil {
		return nil, err
	}
	mutualFriendList := utils.GetMutualFriendList(firstUserFriendList, secondUserFriendList)

	var friendIDs []int
	for _, friend := range mutualFriendList {
		friendIDs = append(friendIDs, friend.ID)
	}
	friends, err := fc.userRepo.GetUserByIDs(friendIDs)
	if err != nil {
		return nil, err
	}
	return friends, nil
}
