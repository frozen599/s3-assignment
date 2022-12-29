package repository

import (
	"github.com/frozen599/s3-assignment/api/internal/config"
	"github.com/frozen599/s3-assignment/api/internal/models"
	"github.com/go-pg/pg/v10"
)

type UserRepository interface {
	GetUserByEmail(email string) (*models.User, error)
	GetUserByIDs(ids []int) ([]models.User, error)
}

type userRepository struct {
}

func NewUserRepository() UserRepository {
	return userRepository{}
}

func (r userRepository) GetUserByEmail(email string) (*models.User, error) {
	var ret models.User
	err := config.GetDB().
		Model(&ret).
		Where("email = ?", email).
		Select()
	if err == pg.ErrNoRows {
		return nil, nil
	}
	return &ret, err
}

func (r userRepository) GetUserByIDs(ids []int) ([]models.User, error) {
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
