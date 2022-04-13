package usecase

import (
	"context"

	"go-backend-template/internal/database"
	"go-backend-template/internal/usecase/dto"
	"go-backend-template/internal/util/crypto"
	"go-backend-template/internal/util/errors"
)

type AuthUsecases interface {
	Login(ctx context.Context, in dto.UserLogin) (dto.UserLoggedInfo, error)
	VerifyAccessToken(accessToken string) (int64, error)
	ParseAccessToken(accessToken string) (int64, error)
}

type authUsecases struct {
	db     database.Service
	config Config
}

func (u *authUsecases) Login(ctx context.Context, in dto.UserLogin) (out dto.UserLoggedInfo, err error) {
	user, err := u.db.UserRepo().GetByEmail(ctx, in.Email)
	if err != nil {
		return out, errors.Wrap(errors.WrongCredentialsError, err, "")
	}
	if !user.ComparePassword(in.Password) {
		return out, errors.New(errors.WrongCredentialsError, "")
	}
	token, err := u.generateAccessToken(user.Id)
	if err != nil {
		return out, err
	}

	return out.MapFrom(user, token), nil
}

func (u *authUsecases) VerifyAccessToken(accessToken string) (int64, error) {
	payload, err := crypto.ParseAndValidateJWT(accessToken, u.config.AccessTokenSecret())
	if err != nil {
		return 0, errors.New(errors.UnauthorizedError, "")
	}

	userId := payload["userId"].(float64)

	return int64(userId), nil
}

func (u *authUsecases) ParseAccessToken(accessToken string) (int64, error) {
	payload, err := crypto.ParseJWT(accessToken, u.config.AccessTokenSecret())
	if err != nil {
		return 0, errors.New(errors.UnauthorizedError, "")
	}

	userId := payload["userId"].(float64)

	return int64(userId), nil
}

func (u *authUsecases) generateAccessToken(userId int64) (string, error) {
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
