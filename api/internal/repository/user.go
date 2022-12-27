package repository

import (
	"github.com/frozen599/s3-assignment/api/internal/config"
	"github.com/frozen599/s3-assignment/api/internal/models"
)

func GetUserByEmail(email string) (*models.User, error) {
	var ret models.User
	err := config.GetDB().
		Model(&ret).
		Where("email = ?", email).
		Select()

	return &ret, err
}
