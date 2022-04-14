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
	userId := int64(1)
	password := "password"
	passwordHash := "password-hash"

	in := dto.UserAdd{
		FirstName: "FirstName",
		LastName:  "LastName",
		Email:     "user@email.com",
		Password:  password,
	}
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

	t.Run("expect it adds new user", func(t *testing.T) {
		prep := newTestPrep()

		prep.crypto.EXPECT().HashPassword(password).Return(passwordHash, nil)
		prep.userRepo.EXPECT().Add(mock.Anything, createUser).Return(userId, nil)
		prep.userRepo.EXPECT().Update(mock.Anything, updateUser).Return(userId, nil)

		actualUserId, err := prep.userUsecases.Add(prep.ctx, in)

		require.NoError(t, err)
		require.Equal(t, userId, actualUserId)
	})

	t.Run("expect it fails if password hashing fails", func(t *testing.T) {
		prep := newTestPrep()
		err := errors.New("password hashing failed")

		prep.crypto.EXPECT().HashPassword(password).Return("", err)

		_, actualErr := prep.userUsecases.Add(prep.ctx, in)

		require.Error(t, actualErr)
		require.EqualError(t, err, actualErr.Error())
	})

	t.Run("expect it fails if user creating fails", func(t *testing.T) {
		prep := newTestPrep()
		err := errors.New("user creating failed")

		prep.crypto.EXPECT().HashPassword(password).Return(passwordHash, nil)
		prep.userRepo.EXPECT().Add(mock.Anything, createUser).Return(userId, err)

		_, actualErr := prep.userUsecases.Add(prep.ctx, in)

		require.Error(t, actualErr)
		require.EqualError(t, err, actualErr.Error())
	})

	t.Run("expect it fails if user updating fails", func(t *testing.T) {
		prep := newTestPrep()
		err := errors.New("user updating failed")

		prep.crypto.EXPECT().HashPassword(password).Return(passwordHash, nil)
		prep.userRepo.EXPECT().Add(mock.Anything, createUser).Return(userId, nil)
		prep.userRepo.EXPECT().Update(mock.Anything, updateUser).Return(userId, err)

		_, actualErr := prep.userUsecases.Add(prep.ctx, in)

		require.Error(t, actualErr)
		require.EqualError(t, err, actualErr.Error())
	})
}
