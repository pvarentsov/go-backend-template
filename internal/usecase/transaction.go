package usecase

import (
	"context"

	"go-backend-template/internal/database"
	"go-backend-template/internal/dto"
)

type TransactionUsecases struct {
	dbService database.Service
	config    Config
}

// AddTwoUsersWithSameEmail Just to show usecase example with transaction
func (u *TransactionUsecases) AddTwoUsersWithSameEmail(ctx context.Context) (int64, error) {
	addUserDTO := dto.AddUser{
		FirstName: "test",
		LastName:  "test",
		Email:     "test@email.com",
		Password:  "qwerty1",
	}

	user, err := addUserDTO.MapTo()
	if err != nil {
		return 0, err
	}

	tx, err := u.dbService.BeginTx(ctx)
	if err != nil {
		return 0, err
	}

	userId, err := tx.Repositories().User.Add(ctx, user)
	if err != nil {
		if err := tx.Rollback(ctx); err != nil {
			return 0, err
		}
		return 0, err
	}

	userId, err = tx.Repositories().User.Add(ctx, user)
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
