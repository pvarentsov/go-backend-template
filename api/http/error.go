package http

import (
	"net/http"

	"go-backend-template/internal/base/errors"
)

func parseError(err error) (status int, message, details string) {
	var baseErr *errors.Error

	if castErr, ok := err.(*errors.Error); ok {
		baseErr = castErr
	}
	if baseErr == nil {
		baseErr = errors.Wrap(err, errors.InternalError, "")
	}

	switch baseErr.Status() {
	case errors.BadRequestError:
		return http.StatusBadRequest, baseErr.Error(), baseErr.DetailedError()
	case errors.ValidationError:
		return http.StatusBadRequest, baseErr.Error(), baseErr.DetailedError()
	case errors.UnauthorizedError:
		return http.StatusUnauthorized, baseErr.Error(), baseErr.DetailedError()
	case errors.WrongCredentialsError:
		return http.StatusUnauthorized, baseErr.Error(), baseErr.DetailedError()
	case errors.NotFoundError:
		return http.StatusNotFound, baseErr.Error(), baseErr.DetailedError()
	case errors.AlreadyExistsError:
		return http.StatusConflict, baseErr.Error(), baseErr.DetailedError()
	default:
		return http.StatusInternalServerError, baseErr.Error(), baseErr.DetailedError()
	}
}
