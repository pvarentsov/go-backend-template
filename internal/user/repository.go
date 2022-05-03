//go:generate mockery --name UserRepository --filename repository.go --output ./mock --with-expecter

package user

import (
	"context"
)

type UserRepository interface {
	Add(ctx context.Context, user UserModel) (int64, error)
	Update(ctx context.Context, user UserModel) (int64, error)
	GetById(ctx context.Context, userId int64) (UserModel, error)
	GetByEmail(ctx context.Context, email string) (UserModel, error)
}
