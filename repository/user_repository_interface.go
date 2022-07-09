package repository

import (
	"context"
	"github.com/risangpratama13/sirka-test/model"
)

type UserRepository interface {
	FindById(ctx context.Context, userid string) (model.User, error)
	FindAll(ctx context.Context) []model.User
}
