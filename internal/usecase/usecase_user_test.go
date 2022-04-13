package usecase

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"go-backend-template/internal/model"
	"go-backend-template/internal/usecase/dto"
)

func TestUserUsecases_Add(t *testing.T) {
	test := newTestUsecases()

	userId := int64(1)
	password := "password"
	passwordHash := "password-hash"

	t.Run("expect it adds new user", func(t *testing.T) {
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

		test.crypto.EXPECT().HashPassword(password).Return(passwordHash, nil)
		test.userRepo.EXPECT().Add(mock.Anything, createUser).Return(userId, nil)
		test.userRepo.EXPECT().Update(mock.Anything, updateUser).Return(userId, nil)

		actualUserId, err := test.userUsecases.Add(test.ctx, in)

		require.NoError(t, err)
		require.Equal(t, userId, actualUserId)
	})
}
