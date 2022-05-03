//go:generate mockery --name Config --filename config.go --output ./mock --with-expecter

package auth

import "time"

type Config interface {
	AccessTokenSecret() string
	AccessTokenExpiresDate() time.Time
}
