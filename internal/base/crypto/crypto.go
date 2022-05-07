//go:generate mockery --name Crypto --filename crypto.go --output ./mock --with-expecter

package crypto

import (
	"time"
)

type Crypto interface {
	HashPassword(password string) (string, error)
	CompareHashAndPassword(hash string, password string) bool

	GenerateJWT(payload map[string]interface{}, secret string, exp time.Time) (string, error)
	ParseAndValidateJWT(token string, secret string) (map[string]interface{}, error)
	ParseJWT(token string, secret string) (map[string]interface{}, error)

	GenerateUUID() (string, error)
}
