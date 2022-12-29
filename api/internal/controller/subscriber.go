package controller

import (
	"time"

	"github.com/frozen599/s3-assignment/api/internal/models"
	"github.com/frozen599/s3-assignment/api/internal/repository"
)

type SubscriberController interface {
	CreateSubScription(requestor, target string) error
}

type subscriberController struct {
	userRepo repository.UserRepository
	relaRepo repository.RelationshipRepository
}

func NewSubscriberController() SubscriberController {
	return subscriberController{
		userRepo: repository.NewUserRepository(),
		relaRepo: repository.NewRelationshipRepository(),
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
