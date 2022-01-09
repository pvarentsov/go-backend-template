package repository

import (
	"go-backend-template/internal/errors"
	"go-backend-template/internal/model"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
)

// User

func parseAddUserError(user *model.User, err error) error {
	pgError, isPgError := err.(*pgconn.PgError)

	if isPgError && pgError.Code == pgerrcode.UniqueViolation {
		switch pgError.ConstraintName {
		case "users_email_key":
			return errors.
				Errorf(errors.AlreadyExistsError, "user with email \"%s\" already exists", user.Email).
				SetInternal(err)
		default:
			return errors.New(errors.DatabaseError, "add user failed").SetInternal(err)
		}
	}

	return errors.New(errors.DatabaseError, "add user failed").SetInternal(err)
}

func parseUpdateUserError(user *model.User, err error) error {
	return errors.New(errors.DatabaseError, "update user failed").SetInternal(err)
}

func parseGetUserByIdError(userId int64, err error) error {
	pgError, isPgError := err.(*pgconn.PgError)

	if isPgError && pgError.Code == pgerrcode.NoDataFound {
		return errors.
			Errorf(errors.NotFoundError, "user with id \"%d\" not found", userId).
			SetInternal(err)
	}
	if err.Error() == "no rows in result set" {
		return errors.
			Errorf(errors.NotFoundError, "user with id \"%d\" not found", userId).
			SetInternal(err)
	}

	return errors.New(errors.DatabaseError, "get user by id failed").SetInternal(err)
}

func parseGetUserByEmailError(email string, err error) error {
	pgError, isPgError := err.(*pgconn.PgError)

	if isPgError && pgError.Code == pgerrcode.NoDataFound {
		return errors.
			Errorf(errors.NotFoundError, "user with email \"%s\" not found", email).
			SetInternal(err)
	}
	if err.Error() == "no rows in result set" {
		return errors.
			Errorf(errors.NotFoundError, "user with email \"%s\" not found", email).
			SetInternal(err)
	}

	return errors.New(errors.DatabaseError, "get user by email failed").SetInternal(err)
}
