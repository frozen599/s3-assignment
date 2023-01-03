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
	userRepo repo.UserRepo
	relaRepo repo.RelationshipRepo
}

func NewBlockingController(userRepo repo.UserRepo, relaRepo repo.RelationshipRepo) BlockingController {
	return blockingController{
		userRepo: userRepo,
		relaRepo: relaRepo,
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
