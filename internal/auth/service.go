//go:generate mockery --name AuthService --filename service.go --output ./mock --with-expecter
//go:generate mockery --name Config --filename config.go --output ./mock --with-expecter

package auth

import (
	"context"
	"time"
)

type AuthService interface {
	Login(ctx context.Context, dto LoginUserDto) (LoggedUserDto, error)
	VerifyAccessToken(accessToken string) (int64, error)
	ParseAccessToken(accessToken string) (int64, error)
}

type Config interface {
	AccessTokenSecret() string
	AccessTokenExpiresDate() time.Time
}
