package impl

import (
	"context"

	"go-backend-template/internal/base/crypto"
	"go-backend-template/internal/base/database"
	"go-backend-template/internal/user"
)

type UserUsecasesOpts struct {
	TxManager      database.TxManager
	UserRepository user.UserRepository
	Crypto         crypto.Crypto
}

func NewUserUsecases(opts UserUsecasesOpts) user.UserUsecases {
	return &userUsecases{
		TxManager:      opts.TxManager,
		UserRepository: opts.UserRepository,
		Crypto:         opts.Crypto,
	}
}

type userUsecases struct {
	database.TxManager
	user.UserRepository
	crypto.Crypto
}

func (u *userUsecases) Add(ctx context.Context, in user.AddUserDto) (userId int64, err error) {
	model, err := in.MapToModel()
	if err != nil {
		return 0, err
	}
	if err := model.HashPassword(u.Crypto); err != nil {
		return 0, err
	}

	// Transaction demonstration
	err = u.RunTx(ctx, func(ctx context.Context) error {
		userId, err = u.UserRepository.Add(ctx, model)
		if err != nil {
			return err
		}
		model.Id = userId

		userId, err = u.UserRepository.Update(ctx, model)
		if err != nil {
			return err
		}
		return nil
	})

	return userId, err
}

func (u *userUsecases) Update(ctx context.Context, in user.UpdateUserDto) (err error) {
	model, err := u.UserRepository.GetById(ctx, in.Id)
	if err != nil {
		return err
	}
	err = model.Update(in.FirstName, in.LastName, in.Email)
	if err != nil {
		return err
	}
	_, err = u.UserRepository.Update(ctx, model)

	return err
}

func (u *userUsecases) ChangePassword(ctx context.Context, in user.ChangeUserPasswordDto) (err error) {
	user, err := u.UserRepository.GetById(ctx, in.Id)
	if err != nil {
		return err
	}
	if err = user.ChangePassword(in.Password, u.Crypto); err != nil {
		return err
	}
	_, err = u.UserRepository.Update(ctx, user)

	return err
}

func (u *userUsecases) GetById(ctx context.Context, userId int64) (out user.UserDto, err error) {
	model, err := u.UserRepository.GetById(ctx, userId)
	if err != nil {
		return out, err
	}

	return out.MapFromModel(model), nil
}
