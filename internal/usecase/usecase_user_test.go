package usecase

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"go-backend-template/internal/model"
	"go-backend-template/internal/usecase/dto"
)

func TestUserUsecases_Add(t *testing.T) {
	prep := newTestPrep()

	userId := int64(1)
	password := "password"
	passwordHash := "password-hash"

	in := dto.UserAdd{
		FirstName: "FirstName",
		LastName:  "LastName",
		Email:     "user@email.com",
		Password:  password,
	}

	t.Run("expect it adds new user", func(t *testing.T) {
		createUser := model.User{
			FirstName: in.FirstName,
			LastName:  in.LastName,
			Email:     in.Email,
			Password:  passwordHash,
		}
		updateUser := model.User{
			Id:        userId,
			FirstName: createUser.FirstName,
			LastName:  createUser.LastName,
			Email:     createUser.Email,
			Password:  createUser.Password,
		}

		prep.crypto.EXPECT().HashPassword(password).
			Times(1).
			Once().
			Return(passwordHash, nil)
		prep.userRepo.EXPECT().Add(mock.Anything, createUser).
			Times(0).
			Once().
			Return(userId, nil)
		prep.userRepo.EXPECT().Update(mock.Anything, updateUser).
			Times(0).
			Once().
			Return(userId, nil)

		actualUserId, err := prep.userUsecases.Add(prep.ctx, in)

		require.NoError(t, err)
		require.Equal(t, userId, actualUserId)
	})

	t.Run("expect it fails if password is weak", func(t *testing.T) {
		err := errors.New("weak password")

		prep.crypto.EXPECT().HashPassword(password).
			Times(1).
			Once().
			Return("", err)
		prep.userRepo.EXPECT().Add(mock.Anything, mock.Anything).
			Times(0).
			Once().
			Return(userId, nil)
		prep.userRepo.EXPECT().Update(mock.Anything, mock.Anything).
			Times(0).
			Once().
			Return(userId, nil)

		_, actualErr := prep.userUsecases.Add(prep.ctx, in)

		require.Error(t, actualErr)
		require.EqualError(t, err, actualErr.Error())
	})
}
