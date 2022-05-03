package impl

import (
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"

	"go-backend-template/internal/base/errors"
	"go-backend-template/internal/user"
)

func parseAddUserError(user *user.UserModel, err error) error {
	pgError, isPgError := err.(*pgconn.PgError)

	if isPgError && pgError.Code == pgerrcode.UniqueViolation {
		switch pgError.ConstraintName {
		case "users_email_key":
			return errors.Wrapf(errors.AlreadyExistsError, err, "user with email \"%s\" already exists", user.Email)
		default:
			return errors.Wrap(errors.DatabaseError, err, "add user failed")
		}
	}

	return errors.Wrap(errors.DatabaseError, err, "add user failed")
}

func parseUpdateUserError(user *user.UserModel, err error) error {
	pgError, isPgError := err.(*pgconn.PgError)

	if isPgError && pgError.Code == pgerrcode.UniqueViolation {
		return errors.Wrapf(errors.AlreadyExistsError, err, "user with email \"%s\" already exists", user.Email)
	}

	return errors.Wrap(errors.DatabaseError, err, "update user failed")
}

func parseGetUserByIdError(userId int64, err error) error {
	pgError, isPgError := err.(*pgconn.PgError)

	if isPgError && pgError.Code == pgerrcode.NoDataFound {
		return errors.Wrapf(errors.NotFoundError, err, "user with id \"%d\" not found", userId)
	}
	if err.Error() == "no rows in result set" {
		return errors.Wrapf(errors.NotFoundError, err, "user with id \"%d\" not found", userId)
	}

	return errors.Wrap(errors.DatabaseError, err, "get user by id failed")
}

func parseGetUserByEmailError(email string, err error) error {
	pgError, isPgError := err.(*pgconn.PgError)

	if isPgError && pgError.Code == pgerrcode.NoDataFound {
		return errors.Wrapf(errors.NotFoundError, err, "user with email \"%s\" not found", email)
	}
	if err.Error() == "no rows in result set" {
		return errors.Wrapf(errors.NotFoundError, err, "user with email \"%s\" not found", email)
	}

	return errors.Wrap(errors.DatabaseError, err, "get user by email failed")
}
