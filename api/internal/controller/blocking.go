package controller

import (
	"time"

	"github.com/frozen599/s3-assignment/api/internal/models"
	"github.com/frozen599/s3-assignment/api/internal/repo"
)

type BlockingController interface {
	BlockUpdate(requestor, target string) error
}

type blockingController struct {
}

func NewBlockingController() BlockingController {
	return blockingController{}
}

func (bc blockingController) BlockUpdate(requestor string, target string) error {
	requestorUser, err := repo.UserRepo.GetUserByEmail(requestor)
	if err != nil {
		return err
	}
	targetUser, err := repo.UserRepo.GetUserByEmail(target)
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

	err = repo.RelationshipRepo.CreateRelationship(blockingRelationShip)
	if err != nil {
		return err
	}
	return nil
}
