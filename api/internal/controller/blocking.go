package controller

import (
	"time"

	"github.com/frozen599/s3-assignment/api/internal/models"
	"github.com/frozen599/s3-assignment/api/internal/repository"
)

type BlockingController interface {
	BlockUpdate(requestor, target string) error
}

type blockingController struct {
	userRepo repository.UserRepository
	relaRepo repository.RelationshipRepository
}

func NewBlockingController() BlockingController {
	return blockingController{
		userRepo: repository.NewUserRepository(),
		relaRepo: repository.NewRelationshipRepository(),
	}
}

func (bc blockingController) BlockUpdate(requestor string, target string) error {
	requestorUser, err := bc.userRepo.GetUserByEmail(requestor)
	if err != nil {
		return err
	}
	targetUser, err := bc.userRepo.GetUserByEmail(target)
	if err != nil {
		return err
	}

	blockingRelationShip := models.Relationship{
		UserID1:          requestorUser.ID,
		UserID2:          targetUser.ID,
		RelationshipType: models.RelationshipTypeBlocking,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	err = bc.relaRepo.CreateRelationship(blockingRelationShip)
	if err != nil {
		return err
	}
	return nil
}
