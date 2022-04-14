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

func TestUserUsecases_UpdateInfo(t *testing.T) {
	in := dto.UserUpdateInfo{
		Id:        int64(2),
		FirstName: "UpdateFirstName",
		LastName:  "UpdateLastName",
		Email:     "user+update@email.com",
	}
	getUser := model.User{
		Id:        in.Id,
		FirstName: "FirstName",
		LastName:  "LastName",
		Email:     "user@email.com",
		Password:  "password-hash",
	}
	updateUser := model.User{
		Id:        in.Id,
		FirstName: in.FirstName,
		LastName:  in.LastName,
		Email:     in.Email,
		Password:  getUser.Password,
	}

	t.Run("expect it updates user", func(t *testing.T) {
		prep := newTestPrep()

		prep.userRepo.EXPECT().GetById(mock.Anything, in.Id).Return(getUser, nil)
		prep.userRepo.EXPECT().Update(mock.Anything, updateUser).Return(in.Id, nil)

		err := prep.userUsecases.UpdateInfo(prep.ctx, in)

		require.NoError(t, err)
	})

	t.Run("expect it fails if user getting fails", func(t *testing.T) {
		prep := newTestPrep()
		err := errors.New("user getting failed")

		prep.userRepo.EXPECT().GetById(mock.Anything, in.Id).Return(getUser, err)

		actualErr := prep.userUsecases.UpdateInfo(prep.ctx, in)

		require.Error(t, actualErr)
		require.EqualError(t, err, actualErr.Error())
	})

	t.Run("expect it fails if user updating fails", func(t *testing.T) {
		prep := newTestPrep()
		err := errors.New("user updating failed")

		prep.userRepo.EXPECT().GetById(mock.Anything, in.Id).Return(getUser, nil)
		prep.userRepo.EXPECT().Update(mock.Anything, updateUser).Return(in.Id, err)

		actualErr := prep.userUsecases.UpdateInfo(prep.ctx, in)

		require.Error(t, actualErr)
		require.EqualError(t, err, actualErr.Error())
	})
}

func TestUserUsecases_ChangePassword(t *testing.T) {
	in := dto.UserChangePassword{
		Id:       int64(3),
		Password: "new-password",
	}
	getUser := model.User{
		Id:        in.Id,
		FirstName: "FirstName",
		LastName:  "LastName",
		Email:     "user@email.com",
		Password:  "old-password-hash",
	}
	updateUser := model.User{
		Id:        getUser.Id,
		FirstName: getUser.FirstName,
		LastName:  getUser.LastName,
		Email:     getUser.Email,
		Password:  "new-password-hash",
	}

	t.Run("expect it changes user password", func(t *testing.T) {
		prep := newTestPrep()

		prep.userRepo.EXPECT().GetById(mock.Anything, in.Id).Return(getUser, nil)
		prep.crypto.EXPECT().HashPassword(in.Password).Return(updateUser.Password, nil)
		prep.userRepo.EXPECT().Update(mock.Anything, updateUser).Return(in.Id, nil)

		err := prep.userUsecases.ChangePassword(prep.ctx, in)

		require.NoError(t, err)
	})

	t.Run("expect it fails if user getting fails", func(t *testing.T) {
		prep := newTestPrep()
		err := errors.New("user getting failed")

		prep.userRepo.EXPECT().GetById(mock.Anything, in.Id).Return(getUser, err)

		actualErr := prep.userUsecases.ChangePassword(prep.ctx, in)

		require.Error(t, actualErr)
		require.EqualError(t, err, actualErr.Error())
	})

	t.Run("expect it fails if password hashing fails", func(t *testing.T) {
		prep := newTestPrep()
		err := errors.New("password hashing failed")

		prep.userRepo.EXPECT().GetById(mock.Anything, in.Id).Return(getUser, nil)
		prep.crypto.EXPECT().HashPassword(in.Password).Return(updateUser.Password, err)

		actualErr := prep.userUsecases.ChangePassword(prep.ctx, in)

		require.Error(t, actualErr)
		require.EqualError(t, err, actualErr.Error())
	})

	t.Run("expect it fails if user updating fails", func(t *testing.T) {
		prep := newTestPrep()
		err := errors.New("user updating failed")

		prep.userRepo.EXPECT().GetById(mock.Anything, in.Id).Return(getUser, nil)
		prep.crypto.EXPECT().HashPassword(in.Password).Return(updateUser.Password, nil)
		prep.userRepo.EXPECT().Update(mock.Anything, updateUser).Return(in.Id, err)

		actualErr := prep.userUsecases.ChangePassword(prep.ctx, in)

		require.Error(t, actualErr)
		require.EqualError(t, err, actualErr.Error())
	})
}

func TestUserUsecases_GetById(t *testing.T) {
	userId := int64(4)

	getUser := model.User{
		Id:        userId,
		FirstName: "FirstName",
		LastName:  "LastName",
		Email:     "user@email.com",
		Password:  "password-hash",
	}
	out := dto.User{
		Id:        getUser.Id,
		FirstName: getUser.FirstName,
		LastName:  getUser.LastName,
		Email:     getUser.Email,
	}

	t.Run("expect it gets user", func(t *testing.T) {
		prep := newTestPrep()

		prep.userRepo.EXPECT().GetById(mock.Anything, userId).Return(getUser, nil)

		actualOut, err := prep.userUsecases.GetById(prep.ctx, userId)

		require.NoError(t, err)
		require.Equal(t, out, actualOut)
	})

	t.Run("expect it fails if user getting fails", func(t *testing.T) {
		prep := newTestPrep()
		err := errors.New("user getting failed")

		prep.userRepo.EXPECT().GetById(mock.Anything, userId).Return(getUser, err)

		_, actualErr := prep.userUsecases.GetById(prep.ctx, userId)

		require.Error(t, actualErr)
		require.EqualError(t, err, actualErr.Error())
	})
}
