package usecase

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"go-backend-template/internal/model"
	"go-backend-template/internal/usecase/dto"
	coreErrors "go-backend-template/internal/util/errors"
)

func TestAuthUsecases_Login(t *testing.T) {
	userId := int64(1)

	token := "token"
	tokenSecret := "token-secret"
	tokenExpires := time.Now().Add(time.Hour)
	tokenPayload := map[string]interface{}{"userId": userId}

	password := "password"
	passwordHash := "password-hash"

	in := dto.UserLogin{
		Email:    "user@email.com",
		Password: password,
	}
	getUser := model.User{
		Id:        userId,
		FirstName: "FirstName",
		LastName:  "LastName",
		Email:     in.Email,
		Password:  passwordHash,
	}
	loginUser := dto.UserLoggedInfo{
		User: dto.User{
			Id:        getUser.Id,
			FirstName: getUser.FirstName,
			LastName:  getUser.LastName,
			Email:     getUser.Email,
		},
		Token: token,
	}

	t.Run("expect it logins user", func(t *testing.T) {
		prep := newTestPrep()

		prep.userRepo.EXPECT().GetByEmail(mock.Anything, in.Email).Return(getUser, nil)
		prep.crypto.EXPECT().CompareHashAndPassword(passwordHash, password).Return(true)

		prep.config.EXPECT().AccessTokenSecret().Return(tokenSecret)
		prep.config.EXPECT().AccessTokenExpiresDate().Return(tokenExpires)
		prep.crypto.EXPECT().GenerateJWT(tokenPayload, tokenSecret, tokenExpires).Return(token, nil)

		actualLoginUser, err := prep.authUsecases.Login(prep.ctx, in)

		require.NoError(t, err)
		require.Equal(t, loginUser, actualLoginUser)
	})

	t.Run("expect it fails if user with such email does't exist", func(t *testing.T) {
		prep := newTestPrep()

		err := errors.New("user not found")
		wrapErr := coreErrors.New(coreErrors.WrongCredentialsError, "")

		prep.userRepo.EXPECT().GetByEmail(mock.Anything, in.Email).Return(getUser, err)

		_, actualErr := prep.authUsecases.Login(prep.ctx, in)

		require.Error(t, actualErr)
		require.EqualError(t, wrapErr, actualErr.Error())
	})

	t.Run("expect it fails if password is wrong", func(t *testing.T) {
		prep := newTestPrep()
		err := coreErrors.New(coreErrors.WrongCredentialsError, "")

		prep.userRepo.EXPECT().GetByEmail(mock.Anything, in.Email).Return(getUser, nil)
		prep.crypto.EXPECT().CompareHashAndPassword(passwordHash, password).Return(false)

		_, actualErr := prep.authUsecases.Login(prep.ctx, in)

		require.Error(t, actualErr)
		require.EqualError(t, err, actualErr.Error())
	})

	t.Run("expect it fails if token generation fails", func(t *testing.T) {
		prep := newTestPrep()
		err := errors.New("token generation failed")

		prep.userRepo.EXPECT().GetByEmail(mock.Anything, in.Email).Return(getUser, nil)
		prep.crypto.EXPECT().CompareHashAndPassword(passwordHash, password).Return(true)

		prep.config.EXPECT().AccessTokenSecret().Return(tokenSecret)
		prep.config.EXPECT().AccessTokenExpiresDate().Return(tokenExpires)
		prep.crypto.EXPECT().GenerateJWT(tokenPayload, tokenSecret, tokenExpires).Return(token, err)

		_, actualErr := prep.authUsecases.Login(prep.ctx, in)

		require.Error(t, actualErr)
		require.EqualError(t, err, actualErr.Error())
	})
}

func TestAuthUsecases_VerifyAccessToken(t *testing.T) {
	userId := int64(1)

	token := "token"
	tokenSecret := "token-secret"
	tokenPayload := map[string]interface{}{"userId": float64(userId)}

	t.Run("expect it virifies token", func(t *testing.T) {
		prep := newTestPrep()

		prep.config.EXPECT().AccessTokenSecret().Return(tokenSecret)
		prep.crypto.EXPECT().ParseAndValidateJWT(token, tokenSecret).Return(tokenPayload, nil)

		actualUserId, err := prep.authUsecases.VerifyAccessToken(token)

		require.NoError(t, err)
		require.Equal(t, userId, actualUserId)
	})

	t.Run("expect it fails if token is not valid", func(t *testing.T) {
		prep := newTestPrep()

		err := errors.New("token is not valid")
		wrapErr := coreErrors.New(coreErrors.UnauthorizedError, "")

		prep.config.EXPECT().AccessTokenSecret().Return(tokenSecret)
		prep.crypto.EXPECT().ParseAndValidateJWT(token, tokenSecret).Return(tokenPayload, err)

		_, actualErr := prep.authUsecases.VerifyAccessToken(token)

		require.Error(t, actualErr)
		require.Equal(t, wrapErr, actualErr)
	})
}

func TestAuthUsecases_ParseAccessToken(t *testing.T) {
	userId := int64(1)

	token := "token"
	tokenSecret := "token-secret"
	tokenPayload := map[string]interface{}{"userId": float64(userId)}

	t.Run("expect it virifies token", func(t *testing.T) {
		prep := newTestPrep()

		prep.config.EXPECT().AccessTokenSecret().Return(tokenSecret)
		prep.crypto.EXPECT().ParseJWT(token, tokenSecret).Return(tokenPayload, nil)

		actualUserId, err := prep.authUsecases.ParseAccessToken(token)

		require.NoError(t, err)
		require.Equal(t, userId, actualUserId)
	})

	t.Run("expect it fails if token parsing fails", func(t *testing.T) {
		prep := newTestPrep()

		err := errors.New("token parsing failed")
		wrapErr := coreErrors.New(coreErrors.UnauthorizedError, "")

		prep.config.EXPECT().AccessTokenSecret().Return(tokenSecret)
		prep.crypto.EXPECT().ParseJWT(token, tokenSecret).Return(tokenPayload, err)

		_, actualErr := prep.authUsecases.ParseAccessToken(token)

		require.Error(t, actualErr)
		require.Equal(t, wrapErr, actualErr)
	})
}
