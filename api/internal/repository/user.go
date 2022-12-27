package repository

import (
	"github.com/frozen599/s3-assignment/api/internal/config"
	"github.com/frozen599/s3-assignment/api/internal/models"
	"github.com/go-pg/pg/v10"
)

func GetUserByEmail(email string) (*models.User, error) {
	var ret models.User
	err := config.GetDB().
		Model(&ret).
		Where("email = ?", email).
		Select()

	return &ret, err
}

func GetUserByIds(ids []int) ([]models.User, error) {
	var ret []models.User

	err := config.GetDB().
		Model(&ret).
		Where("id IN (?)", pg.In(ids)).
		Select()
	if err == pg.ErrNoRows {
		return nil, nil
	}
	return ret, nil
}
