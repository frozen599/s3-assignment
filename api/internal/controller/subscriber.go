package controller

import (
	"time"

	"github.com/frozen599/s3-assignment/api/internal/models"
	"github.com/frozen599/s3-assignment/api/internal/pkg"
	"github.com/frozen599/s3-assignment/api/internal/repo"
)

type SubscriberController interface {
	CreateSubScription(requestor, target string) error
	GetCanReceiveUpdate(sender string, mentionedEmails []string) ([]string, error)
}

type subscriberController struct {
	userRepo repo.UserRepo
	relaRepo repo.RelationshipRepo
}

func NewSubscriberController(userRepo repo.UserRepo, relaRepo repo.RelationshipRepo) SubscriberController {
	return subscriberController{
		userRepo: userRepo,
		relaRepo: relaRepo,
	}
}

func (sc subscriberController) CreateSubScription(requestor, target string) error {
	requestorUser, err := sc.userRepo.GetUserByEmail(requestor)
	if err != nil {
		return err
	}
	targetUser, err := sc.userRepo.GetUserByEmail(target)
	if err != nil {
		return err
	}
	if requestorUser == nil || targetUser == nil {
		return pkg.ErrUserNotFound
	}

	isAlreadySubscribing, err := sc.relaRepo.CheckIfAlreadySubscribing(requestorUser.ID, targetUser.ID)
	if err != nil {
		return err
	}
	if isAlreadySubscribing {
		return pkg.ErrCurrentUserIsAlreadySubscribingTarget
	}

	isBlockingOrBlocked, err := sc.relaRepo.CheckIfIsBlockingOrBlocked(requestorUser.ID, targetUser.ID)
	if err != nil {
		return err
	}
	if isBlockingOrBlocked {
		return pkg.ErrCurrentUserIsBlockingTargetOrBlocked
	}

	blockingRelationShip := models.Relationship{
		UserID1:          requestorUser.ID,
		UserID2:          targetUser.ID,
		RelationshipType: models.RelationshipTypeSubscriber,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	err = sc.relaRepo.CreateRelationship(blockingRelationShip)
	if err != nil {
		return err
	}
	return nil

}

func (sc subscriberController) GetCanReceiveUpdate(sender string, mentionedEmails []string) ([]string, error) {
	senderUser, err := sc.userRepo.GetUserByEmail(sender)
	if err != nil {
		return nil, err
	}

	var mentionedIDs []int
	if len(mentionedEmails) > 0 {
		mentionedUsers, err := sc.userRepo.GetUserByEmails(mentionedEmails)
		if err != nil {
			return nil, err
		}
		for _, mentionedUser := range mentionedUsers {
			mentionedIDs = append(mentionedIDs, mentionedUser.ID)
		}
	}

	if senderUser != nil {
		followerIDs, err := sc.relaRepo.GetFollowers(senderUser.ID)
		if err != nil {
			return nil, err
		}
		friendIDs, err := sc.relaRepo.GetFriendList(senderUser.ID)
		if err != nil {
			return nil, err
		}
		blockingOrBlockedList, err := sc.relaRepo.GetBlockedOrBlockingList(senderUser.ID, mentionedIDs)
		if err != nil {
			return nil, err
		}
		notBlockingOrBlockedList := pkg.FindDiff(mentionedIDs, blockingOrBlockedList)

		var tempIDs []int
		tempIDs = append(tempIDs, friendIDs...)
		tempIDs = append(tempIDs, followerIDs...)
		tempIDs = append(tempIDs, notBlockingOrBlockedList...)
		userIDs := pkg.RemoveDuplicateIDs(tempIDs)
		users, err := sc.userRepo.GetUserByIDs(userIDs)
		if err != nil {
			return nil, err
		}

		var emails []string
		for _, user := range users {
			emails = append(emails, user.Email)
		}
		return emails, nil
	}
	return nil, pkg.ErrUserNotFound
}
