package repository

import (
	"time"

	"github.com/frozen599/s3-assignment/api/internal/config"
	"github.com/frozen599/s3-assignment/api/internal/forms"
	"github.com/frozen599/s3-assignment/api/internal/models"
)

func CreateSubscription(form *forms.SubscribeToEmailRequest) error {
	requestor, err := GetUserByEmail(form.Requestor)
	if err != nil {
		return err
	}
	target, err := GetUserByEmail(form.Target)
	if err != nil {
		return err
	}

	friendRelationShip := models.Relationship{
		UserID1:          requestor.ID,
		UserID2:          target.ID,
		RelationshipType: models.RelationshipTypeFriend,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	_, err = config.GetDB().
		Model(&friendRelationShip).
		Insert()

	if err != nil {
		return err
	}
	return nil
}
