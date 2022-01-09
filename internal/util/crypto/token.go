package crypto

import (
	"time"

	"go-backend-template/internal/util/errors"

	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt"
)

func GenerateJWT(payload map[string]interface{}, secret string, exp time.Time) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = exp.Unix()

	for key, value := range payload {
		claims[key] = value
	}

	token, err := jwt.
		NewWithClaims(jwt.SigningMethodHS256, claims).
		SignedString([]byte(secret))

	if err != nil {
		return "", err
	}

	return token, nil
}

func ParseAndValidateJWT(token string, secret string) (map[string]interface{}, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return map[string]interface{}{}, err
	}
	if !parsedToken.Valid {
		return map[string]interface{}{}, errors.New(errors.InternalError, "token validation error")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return map[string]interface{}{}, errors.New(errors.InternalError, "token validation error")
	}

	payload := make(map[string]interface{})

	for key, value := range claims {
		payload[key] = value
	}

	return payload, nil
}

func ParseJWT(token string, secret string) (map[string]interface{}, error) {
	parsedToken, _ := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return map[string]interface{}{}, errors.New(errors.InternalError, "token parsing error")
	}

	payload := make(map[string]interface{})

	for key, value := range claims {
		payload[key] = value
	}

	return payload, nil
}

func GenerateUUID() (string, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return "", err
	}

	return id.String(), nil
}
