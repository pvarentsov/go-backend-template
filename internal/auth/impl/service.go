package impl

import (
	"context"

	"go-backend-template/internal/auth"
	"go-backend-template/internal/base/crypto"
	"go-backend-template/internal/base/errors"
	"go-backend-template/internal/user"
)

type AuthServiceOpts struct {
	UserRepository user.UserRepository
	Crypto         crypto.Crypto
	Config         auth.Config
}

func NewAuthService(opts AuthServiceOpts) auth.AuthService {
	return &authService{
		UserRepository: opts.UserRepository,
		Crypto:         opts.Crypto,
		Config:         opts.Config,
	}
}

type authService struct {
	user.UserRepository
	crypto.Crypto
	auth.Config
}

func (u *authService) Login(ctx context.Context, in auth.LoginUserDto) (out auth.LoggedUserDto, err error) {
	user, err := u.UserRepository.GetByEmail(ctx, in.Email)
	if err != nil {
		return out, errors.Wrap(err, errors.WrongCredentialsError, "")
	}
	if !user.ComparePassword(in.Password, u.Crypto) {
		return out, errors.New(errors.WrongCredentialsError, "")
	}
	token, err := u.generateAccessToken(user.Id)
	if err != nil {
		return out, err
	}

	return out.MapFromModel(user, token), nil
}

func (u *authService) VerifyAccessToken(accessToken string) (int64, error) {
	payload, err := u.ParseAndValidateJWT(accessToken, u.AccessTokenSecret())
	if err != nil {
		return 0, errors.New(errors.UnauthorizedError, "")
	}

	userId, ok := payload["userId"].(float64)
	if !ok {
		return 0, errors.New(errors.UnauthorizedError, "")
	}

	return int64(userId), nil
}

func (u *authService) ParseAccessToken(accessToken string) (int64, error) {
	payload, err := u.ParseJWT(accessToken, u.AccessTokenSecret())
	if err != nil {
		return 0, errors.New(errors.UnauthorizedError, "")
	}

	userId, ok := payload["userId"].(float64)
	if !ok {
		return 0, errors.New(errors.UnauthorizedError, "")
	}

	return int64(userId), nil
}

func (u *authService) generateAccessToken(userId int64) (string, error) {
	payload := map[string]interface{}{"userId": userId}

	return u.GenerateJWT(
		payload,
		u.AccessTokenSecret(),
		u.AccessTokenExpiresDate(),
	)
}
