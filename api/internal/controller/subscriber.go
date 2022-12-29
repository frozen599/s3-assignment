package controller

import (
	"time"

	"github.com/frozen599/s3-assignment/api/internal/models"
	"github.com/frozen599/s3-assignment/api/internal/repo"
	"github.com/frozen599/s3-assignment/api/internal/utils"
)

type SubscriberController interface {
	CreateSubScription(requestor, target string) error
	CanReceiveUpdate(sender, text string) ([]string, error)
}

type subscriberController struct {
}

func NewSubscriberController() SubscriberController {
	return subscriberController{}
}

func (sc subscriberController) CreateSubScription(requestor, target string) error {
	requestorUser, err := repo.UserRepo.GetUserByEmail(requestor)
	if err != nil {
		return err
	}
	targetUser, err := repo.UserRepo.GetUserByEmail(target)
	if err != nil {
		return err
	}

	if requestorUser != nil && targetUser != nil {
		blockingRelationShip := models.Relationship{
			UserID1:          requestorUser.ID,
			UserID2:          targetUser.ID,
			RelationshipType: models.RelationshipTypeSubscriber,
			CreatedAt:        time.Now(),
			UpdatedAt:        time.Now(),
		}

		err = repo.RelationshipRepo.CreateRelationship(blockingRelationShip)
		if err != nil {
			return err
		}
		return nil
	}
	return utils.ErrUserNotFound
}

func (sc subscriberController) CanReceiveUpdate(sender, text string) ([]string, error) {
	senderUser, err := repo.UserRepo.GetUserByEmail(sender)
	if err != nil {
		return nil, err
	}

	mentionedEmail := utils.ParseEmail(text)[0]
	mentionedUser, err := repo.UserRepo.GetUserByEmail(mentionedEmail)
	if err != nil {
		return nil, err
	}

	if senderUser != nil {
		relas, err := repo.RelationshipRepo.CanReceiveUpdate(senderUser.ID)
		if err != nil {
			return nil, err
		}
		var userIDs []int
		for _, rela := range relas {
			userIDs = append(userIDs, rela.UserID2)
		}
		users, err := repo.UserRepo.GetUserByIDs(userIDs)
		if err != nil {
			return nil, err
		}
		var emails []string
		for _, user := range users {
			emails = append(emails, user.Email)
		}
		if mentionedUser != nil {
			emails = append(emails, mentionedUser.Email)
		}
		return emails, nil
	}
	return nil, utils.ErrUserNotFound
}
