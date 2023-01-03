package repo

import (
	"github.com/frozen599/s3-assignment/api/internal/models"
	"github.com/go-pg/pg/v10"
)

type UserRepo interface {
	GetUserByEmail(email string) (*models.User, error)
	GetUserByIDs(ids []int) ([]models.User, error)
}

type userRepo struct {
	db *pg.DB
}

func NewUserRepo(dbInstance *pg.DB) UserRepo {
	return userRepo{db: dbInstance}
}

func (r userRepo) GetUserByEmail(email string) (*models.User, error) {
	var ret models.User
	err := r.db.
		Model(&ret).
		Where("email = ?", email).
		Select()
	if err == pg.ErrNoRows {
		return nil, nil
	}
	return &ret, err
}

func (r userRepo) GetUserByIDs(ids []int) ([]models.User, error) {
	var ret []models.User

	err := r.db.
		Model(&ret).
		Where("id IN (?)", pg.In(ids)).
		Select()
	if err == pg.ErrNoRows {
		return nil, nil
	}
	return ret, nil
}
