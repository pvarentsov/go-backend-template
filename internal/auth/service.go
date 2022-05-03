//go:generate mockery --name AuthService --filename service.go --output ./mock --with-expecter

package auth

import (
	"context"
)

type AuthService interface {
	Login(ctx context.Context, dto LoginUserDto) (LoggedUserDto, error)
	VerifyAccessToken(accessToken string) (int64, error)
	ParseAccessToken(accessToken string) (int64, error)
}
