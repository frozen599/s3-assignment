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
}

func NewSubscriberController() SubscriberController {
	return subscriberController{}
}

func (s subscriberController) CreateSubScription(requestor, target string) error {
	requestorUser, err := repository.GetUserByEmail(requestor)
	if err != nil {
		return err
	}
	targetUser, err := repository.GetUserByEmail(target)
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

	err = repository.CreateRelationship(blockingRelationShip)
	if err != nil {
		return err
	}
	return nil
}
