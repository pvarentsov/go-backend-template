package database

import (
	"context"

	"go-backend-template/internal/model"
	"go-backend-template/internal/util/errors"

	"github.com/doug-martin/goqu/v9"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
)

// UserRepo

type UserRepo interface {
	Add(ctx context.Context, user model.User) (int64, error)
	Update(ctx context.Context, user model.User) (int64, error)
	GetById(ctx context.Context, userId int64) (model.User, error)
	GetByEmail(ctx context.Context, email string) (model.User, error)
}

func newUserRepo(pool connection, qb goqu.DialectWrapper) UserRepo {
	return &userRepo{
		repo: repo{pool: pool, qb: qb},
	}
}

type userRepo struct {
	repo
}

func (r *userRepo) Add(ctx context.Context, user model.User) (int64, error) {
	sql, _, err := r.qb.
		Insert("users").
		Rows(goqu.Record{
			"firstname": user.FirstName,
			"lastname":  user.LastName,
			"email":     user.Email,
			"password":  user.Password,
		}).
		Returning("user_id").
		ToSQL()

	if err != nil {
		return 0, errors.New(errors.DatabaseError, "").SetInternal(err)
	}

	row := r.conn(ctx).QueryRow(ctx, sql)

	if err := row.Scan(&user.Id); err != nil {
		return 0, parseAddUserError(&user, err)
	}

	return user.Id, nil
}

func (r *userRepo) Update(ctx context.Context, user model.User) (int64, error) {
	sql, _, err := r.qb.
		Update("users").
		Set(goqu.Record{
			"firstname": user.FirstName,
			"lastname":  user.LastName,
			"email":     user.Email,
			"password":  user.Password,
		}).
		Where(goqu.Ex{"user_id": user.Id}).
		Returning("user_id").
		ToSQL()

	if err != nil {
		return 0, errors.New(errors.DatabaseError, "").SetInternal(err)
	}

	row := r.conn(ctx).QueryRow(ctx, sql)

	if err := row.Scan(&user.Id); err != nil {
		return 0, parseUpdateUserError(&user, err)
	}

	return user.Id, nil
}

func (r *userRepo) GetById(ctx context.Context, userId int64) (model.User, error) {
	sql, _, err := r.qb.
		Select(
			"firstname",
			"lastname",
			"email",
			"password",
		).
		From("users").
		Where(goqu.Ex{"user_id": userId}).
		ToSQL()

	if err != nil {
		return model.User{}, errors.New(errors.DatabaseError, "").SetInternal(err)
	}

	row := r.conn(ctx).QueryRow(ctx, sql)

	user := model.User{Id: userId}

	err = row.Scan(
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
	)
	if err != nil {
		return model.User{}, parseGetUserByIdError(userId, err)
	}

	return user, nil
}

func (r *userRepo) GetByEmail(ctx context.Context, email string) (model.User, error) {
	sql, _, err := r.qb.
		Select(
			"user_id",
			"firstname",
			"lastname",
			"password",
		).
		From("users").
		Where(goqu.Ex{"email": email}).
		ToSQL()

	if err != nil {
		return model.User{}, errors.New(errors.DatabaseError, "").SetInternal(err)
	}

	row := r.conn(ctx).QueryRow(ctx, sql)

	user := model.User{Email: email}

	err = row.Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Password,
	)
	if err != nil {
		return model.User{}, parseGetUserByEmailError(email, err)
	}

	return user, nil
}

// Errors

func parseAddUserError(user *model.User, err error) error {
	pgError, isPgError := err.(*pgconn.PgError)

	if isPgError && pgError.Code == pgerrcode.UniqueViolation {
		switch pgError.ConstraintName {
		case "users_email_key":
			return errors.Errorf(errors.AlreadyExistsError, "user with email \"%s\" already exists", user.Email).
				SetInternal(err)
		default:
			return errors.New(errors.DatabaseError, "add user failed").SetInternal(err)
		}
	}

	return errors.New(errors.DatabaseError, "add user failed").SetInternal(err)
}

func parseUpdateUserError(user *model.User, err error) error {
	pgError, isPgError := err.(*pgconn.PgError)

	if isPgError && pgError.Code == pgerrcode.UniqueViolation {
		return errors.
			Errorf(errors.AlreadyExistsError, "user with email \"%s\" already exists", user.Email).
			SetInternal(err)
	}

	return errors.New(errors.DatabaseError, "update user failed").SetInternal(err)
}

func parseGetUserByIdError(userId int64, err error) error {
	pgError, isPgError := err.(*pgconn.PgError)

	if isPgError && pgError.Code == pgerrcode.NoDataFound {
		return errors.Errorf(errors.NotFoundError, "user with id \"%d\" not found", userId).
			SetInternal(err)
	}
	if err.Error() == "no rows in result set" {
		return errors.Errorf(errors.NotFoundError, "user with id \"%d\" not found", userId).
			SetInternal(err)
	}

	return errors.New(errors.DatabaseError, "get user by id failed").SetInternal(err)
}

func parseGetUserByEmailError(email string, err error) error {
	pgError, isPgError := err.(*pgconn.PgError)

	if isPgError && pgError.Code == pgerrcode.NoDataFound {
		return errors.Errorf(errors.NotFoundError, "user with email \"%s\" not found", email).
			SetInternal(err)
	}
	if err.Error() == "no rows in result set" {
		return errors.Errorf(errors.NotFoundError, "user with email \"%s\" not found", email).
			SetInternal(err)
	}

	return errors.New(errors.DatabaseError, "get user by email failed").SetInternal(err)
}
