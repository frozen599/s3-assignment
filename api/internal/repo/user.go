package repo

import (
	"github.com/frozen599/s3-assignment/api/internal/config"
	"github.com/frozen599/s3-assignment/api/internal/models"
	"github.com/go-pg/pg/v10"
)

type IUserRepo interface {
	GetUserByEmail(email string) (*models.User, error)
	GetUserByIDs(ids []int) ([]models.User, error)
}

type userRepo struct {
}

var UserRepo IUserRepo = userRepo{}

func (userRepo) GetUserByEmail(email string) (*models.User, error) {
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

func (userRepo) GetUserByIDs(ids []int) ([]models.User, error) {
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
