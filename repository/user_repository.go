package repository

import (
	"context"
	"errors"
	"github.com/risangpratama13/sirka-test/model"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	Db *gorm.DB
}

func NewUserRepository(Db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{Db}
}

func (u *UserRepositoryImpl) FindById(ctx context.Context, userid string) (model.User, error) {
	var user model.User
	err := u.Db.Where("userid = ?", userid).First(&user).Error
	if err != nil {
		return user, errors.New("no matches user based on this userid")
	}
	return user, nil
}

func (u *UserRepositoryImpl) FindAll(ctx context.Context) []model.User {
	var users []model.User
	u.Db.Find(&users)

	return users
}
