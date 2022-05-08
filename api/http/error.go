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

	status = convertErrorStatusToHTTP(baseErr.Status())
	message = baseErr.Error()
	details = baseErr.DetailedError()

	return
}

func convertErrorStatusToHTTP(status errors.Status) int {
	switch status {
	case errors.BadRequestError:
		return http.StatusBadRequest
	case errors.ValidationError:
		return http.StatusBadRequest
	case errors.UnauthorizedError:
		return http.StatusUnauthorized
	case errors.WrongCredentialsError:
		return http.StatusUnauthorized
	case errors.NotFoundError:
		return http.StatusNotFound
	case errors.AlreadyExistsError:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
