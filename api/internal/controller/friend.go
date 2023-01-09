package controller

import (
	"time"

	"github.com/frozen599/s3-assignment/api/internal/models"
	"github.com/frozen599/s3-assignment/api/internal/pkg"
	"github.com/frozen599/s3-assignment/api/internal/repo"
)

type FriendController interface {
	CreateFriendConnection(firsrUserEmail, secondUserEmail string) error
	GetFriendList(email string) ([]models.User, error)
	GetMutualFriendList(firsrUserEmail, secondUserEmail string) ([]models.User, error)
}

type friendController struct {
	userRepo repo.UserRepo
	relaRepo repo.RelationshipRepo
}

func NewFriendController(userRepo repo.UserRepo, relaRepo repo.RelationshipRepo) FriendController {
	return friendController{
		userRepo: userRepo,
		relaRepo: relaRepo,
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
	if firstUser == nil || secondUser == nil {
		return pkg.ErrUserNotFound
	}

	existingFriendship, err := fc.relaRepo.CheckIfFriendConnectionExists(firstUser.ID, secondUser.ID)
	if err != nil {
		return err
	}
	if existingFriendship {
		return pkg.ErrFriendshipAlreadyExists
	}

	isBlockingTarget, err := fc.relaRepo.CheckIfIsBlockingTarget(firstUser.ID, secondUser.ID)
	if err != nil {
		return err
	}
	if isBlockingTarget {
		return pkg.ErrCurrentUserIsBlockingTarget
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
	if user == nil {
		return nil, pkg.ErrUserNotFound
	}

	friendListRelationships, err := fc.relaRepo.GetFriendList(user.ID)
	if err != nil {
		return nil, err
	}

	var friendIDs []int
	for _, friend := range friendListRelationships {
		friendIDs = append(friendIDs, friend.UserID1)
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
	if firstUser == nil || secondUser == nil {
		return nil, pkg.ErrUserNotFound
	}

	firstUserFriendList, err := fc.relaRepo.GetFriendList(firstUser.ID)
	if err != nil {
		return nil, err
	}
	secondUserFriendList, err := fc.relaRepo.GetFriendList(secondUser.ID)
	if err != nil {
		return nil, err
	}
	mutualFriendListIDs := pkg.GetMutualFriendList(firstUserFriendList, secondUserFriendList)
	friends, err := fc.userRepo.GetUserByIDs(mutualFriendListIDs)
	if err != nil {
		return nil, err
	}
	return friends, nil
}
