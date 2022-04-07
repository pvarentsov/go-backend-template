package usecase

import (
	"context"

	"go-backend-template/internal/database"
	"go-backend-template/internal/dto"
)

type UserUsecases struct {
	db     database.Service
	config Config
}

func (u *UserUsecases) Add(ctx context.Context, addUserDTO dto.AddUser) (int64, error) {
	user, err := addUserDTO.MapTo()
	if err != nil {
		return 0, err
	}

	// Transaction demonstration

	tx, err := u.db.BeginTx(ctx)
	if err != nil {
		return 0, err
	}

	userId, err := tx.UserRepo.Add(ctx, user)
	if err != nil {
		if err := tx.Rollback(ctx); err != nil {
			return 0, err
		}
		return 0, err
	}

	user.Id = userId

	_, err = tx.UserRepo.Update(ctx, user)
	if err != nil {
		if err := tx.Rollback(ctx); err != nil {
			return 0, err
		}
		return 0, err
	}

	if err := tx.Commit(ctx); err != nil {
		return 0, err
	}

	return userId, nil
}

func (u *UserUsecases) UpdateInfo(ctx context.Context, updateUserInfoDTO dto.UpdateUserInfo) error {
	user, err := u.db.UserRepo.GetById(ctx, updateUserInfoDTO.Id)
	if err != nil {
		return err
	}

	err = user.UpdateInfo(
		updateUserInfoDTO.FirstName,
		updateUserInfoDTO.LastName,
		updateUserInfoDTO.Email,
	)
	if err != nil {
		return err
	}

	_, err = u.db.UserRepo.Update(ctx, user)

	return err
}

func (u *UserUsecases) ChangePassword(ctx context.Context, changeUserPasswordDTO dto.ChangeUserPassword) error {
	user, err := u.db.UserRepo.GetById(ctx, changeUserPasswordDTO.Id)
	if err != nil {
		return err
	}

	if err = user.ChangePassword(changeUserPasswordDTO.Password); err != nil {
		return err
	}

	_, err = u.db.UserRepo.Update(ctx, user)

	return err
}

func (u *UserUsecases) GetById(ctx context.Context, userId int64) (dto.User, error) {
	user, err := u.db.UserRepo.GetById(ctx, userId)
	if err != nil {
		return dto.User{}, err
	}

	userDTO := dto.User{}
	userDTO.MapFrom(user)

	return userDTO, nil
}
