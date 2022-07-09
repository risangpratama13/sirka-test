package service

import (
	"context"
	"github.com/risangpratama13/sirka-test/model"
	"github.com/risangpratama13/sirka-test/repository"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &UserServiceImpl{userRepository}
}

func (service *UserServiceImpl) FindById(ctx context.Context, userid string) (model.User, error) {
	user, err := service.UserRepository.FindById(ctx, userid)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (service *UserServiceImpl) FindAll(ctx context.Context) []model.User {
	users := service.UserRepository.FindAll(ctx)
	return users
}
