package repo

import (
	"context"
	"mall/internal/entity"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user entity.User) (rowId int, err error)
	GetUserByAccount(ctx context.Context, account string, password string) (user entity.User, err error)
	FindUserById(ctx context.Context, id int) (user entity.User, err error)
}
