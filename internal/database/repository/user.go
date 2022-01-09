package repository

import (
	"context"

	"go-backend-template/internal/errors"
	"go-backend-template/internal/model"

	"github.com/doug-martin/goqu/v9"
)

type UserRepository interface {
	Add(ctx context.Context, user model.User) (int64, error)
	Update(ctx context.Context, user model.User) (int64, error)
	GetById(ctx context.Context, userId int64) (model.User, error)
	GetByEmail(ctx context.Context, email string) (model.User, error)
}

func NewUserRepository(queries Queries) UserRepository {
	return &userRepository{
		queries: queries,
	}
}

type userRepository struct {
	queries Queries
}

func (r *userRepository) Add(ctx context.Context, user model.User) (int64, error) {
	sql, _, err := r.queries.Builder.
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

	row, err := r.queries.Row(ctx, sql)
	if err != nil {
		return 0, errors.New(errors.DatabaseError, "").SetInternal(err)
	}

	if err := row.Scan(&user.Id); err != nil {
		return 0, parseAddUserError(&user, err)
	}

	return user.Id, nil
}

func (r *userRepository) Update(ctx context.Context, user model.User) (int64, error) {
	sql, _, err := r.queries.Builder.
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

	row, err := r.queries.Row(ctx, sql)
	if err != nil {
		return 0, errors.New(errors.DatabaseError, "").SetInternal(err)
	}

	if err := row.Scan(&user.Id); err != nil {
		return 0, parseUpdateUserError(&user, err)
	}

	return user.Id, nil
}

func (r *userRepository) GetById(ctx context.Context, userId int64) (model.User, error) {
	sql, _, err := r.queries.Builder.
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

	row, err := r.queries.Row(ctx, sql)
	if err != nil {
		return model.User{}, errors.New(errors.DatabaseError, "").SetInternal(err)
	}

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

func (r *userRepository) GetByEmail(ctx context.Context, email string) (model.User, error) {
	sql, _, err := r.queries.Builder.
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

	row, err := r.queries.Row(ctx, sql)
	if err != nil {
		return model.User{}, errors.New(errors.DatabaseError, "").SetInternal(err)
	}

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
