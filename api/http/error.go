package http

import (
	"net/http"

	"go-backend-template/internal/base/errors"
)

func parseError(err error) (int, string) {
	withDetails := false

	getMessage := func(baseErr errors.Error) string {
		if withDetails {
			return baseErr.ErrorWithDetails()
		}
		return baseErr.Error()
	}

	if baseErr, ok := err.(errors.Error); ok {
		switch baseErr.Status() {
		case errors.BadRequestError:
			return http.StatusBadRequest, getMessage(baseErr)
		case errors.ValidationError:
			return http.StatusBadRequest, getMessage(baseErr)
		case errors.UnauthorizedError:
			return http.StatusUnauthorized, getMessage(baseErr)
		case errors.WrongCredentialsError:
			return http.StatusUnauthorized, getMessage(baseErr)
		case errors.NotFoundError:
			return http.StatusNotFound, getMessage(baseErr)
		case errors.AlreadyExistsError:
			return http.StatusConflict, getMessage(baseErr)
		default:
			return http.StatusInternalServerError, getMessage(baseErr)
		}
	}

	baseErr := errors.Wrap(err, errors.InternalError, "")

	return http.StatusInternalServerError, getMessage(baseErr)
}
