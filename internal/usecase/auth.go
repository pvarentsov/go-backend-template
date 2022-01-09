package usecase

import (
	"context"

	"go-backend-template/internal/database"
	"go-backend-template/internal/dto"
	"go-backend-template/internal/errors"
	"go-backend-template/internal/util/crypto"
)

type AuthUsecases struct {
	dbService database.Service
	config    Config
}

func (u *AuthUsecases) Login(ctx context.Context, loginUserDTO dto.LoginUser) (dto.LoggedUser, error) {
	user, err := u.dbService.Repositories().User.GetByEmail(ctx, loginUserDTO.Email)
	if err != nil {
		return dto.LoggedUser{}, errors.
			New(errors.WrongCredentialsError, "").
			SetInternal(err)
	}

	if !user.ComparePassword(loginUserDTO.Password) {
		return dto.LoggedUser{}, errors.New(errors.WrongCredentialsError, "")
	}

	token, err := u.generateAccessToken(user.Id)
	if err != nil {
		return dto.LoggedUser{}, err
	}

	loggedUserDTO := dto.LoggedUser{}
	loggedUserDTO.MapFrom(user, token)

	return loggedUserDTO, nil
}

func (u *AuthUsecases) VerifyAccessToken(accessToken string) (int64, error) {
	payload, err := crypto.ParseAndValidateJWT(accessToken, u.config.AccessTokenSecret())
	if err != nil {
		return 0, errors.New(errors.UnauthorizedError, "")
	}

	userId := payload["userId"].(float64)

	return int64(userId), nil
}

func (u *AuthUsecases) ParseAccessToken(accessToken string) (int64, error) {
	payload, err := crypto.ParseJWT(accessToken, u.config.AccessTokenSecret())
	if err != nil {
		return 0, errors.New(errors.UnauthorizedError, "")
	}

	userId := payload["userId"].(float64)

	return int64(userId), nil
}

func (u *AuthUsecases) generateAccessToken(userId int64) (string, error) {
	payload := map[string]interface{}{"userId": userId}

	token, err := crypto.GenerateJWT(
		payload,
		u.config.AccessTokenSecret(),
		u.config.AccessTokenExpiresDate(),
	)
	if err != nil {
		return "", err
	}

	return token, nil
}
