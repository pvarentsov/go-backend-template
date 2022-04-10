package http

import (
	"net/http"

	"go-backend-template/internal/util/errors"
)

func parseError(err error) (int, string) {
	if err == nil {
		return http.StatusInternalServerError, "internal error"
	}

	wrappedErr := errors.Wrap(errors.InternalError, err, "internal error")

	switch wrappedErr.Status() {
	case errors.BadRequestError:
		return http.StatusBadRequest, wrappedErr.Error()
	case errors.ValidationError:
		return http.StatusBadRequest, wrappedErr.Error()
	case errors.UnauthorizedError:
		return http.StatusUnauthorized, wrappedErr.Error()
	case errors.WrongCredentialsError:
		return http.StatusUnauthorized, wrappedErr.Error()
	case errors.NotFoundError:
		return http.StatusNotFound, wrappedErr.Error()
	case errors.AlreadyExistsError:
		return http.StatusConflict, wrappedErr.Error()
	default:
		return http.StatusInternalServerError, wrappedErr.Error()
	}
}
