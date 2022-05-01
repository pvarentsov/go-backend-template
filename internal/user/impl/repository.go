package impl

import (
	"context"

	"github.com/doug-martin/goqu/v9"

	"go-backend-template/internal/base/database"
	"go-backend-template/internal/base/errors"
	"go-backend-template/internal/user"
)

type UserRepositoryOpts struct {
	ConnManager database.ConnManager
}

func NewUserRepository(opts UserRepositoryOpts) user.UserRepository {
	return &userRepository{
		ConnManager: opts.ConnManager,
	}
}

type userRepository struct {
	database.ConnManager
}

func (r *userRepository) Add(ctx context.Context, model user.UserModel) (int64, error) {
	sql, _, err := database.QueryBuilder.
		Insert("users").
		Rows(goqu.Record{
			"firstname": model.FirstName,
			"lastname":  model.LastName,
			"email":     model.Email,
			"password":  model.Password,
		}).
		Returning("user_id").
		ToSQL()

	if err != nil {
		return 0, errors.Wrap(errors.DatabaseError, err, "syntax error")
	}

	row := r.Conn(ctx).QueryRow(ctx, sql)

	if err := row.Scan(&model.Id); err != nil {
		return 0, parseAddUserError(&model, err)
	}

	return model.Id, nil
}

func (r *userRepository) Update(ctx context.Context, model user.UserModel) (int64, error) {
	sql, _, err := database.QueryBuilder.
		Update("users").
		Set(goqu.Record{
			"firstname": model.FirstName,
			"lastname":  model.LastName,
			"email":     model.Email,
			"password":  model.Password,
		}).
		Where(goqu.Ex{"user_id": model.Id}).
		Returning("user_id").
		ToSQL()

	if err != nil {
		return 0, errors.Wrap(errors.DatabaseError, err, "syntax error")
	}

	row := r.Conn(ctx).QueryRow(ctx, sql)

	if err := row.Scan(&model.Id); err != nil {
		return 0, parseUpdateUserError(&model, err)
	}

	return model.Id, nil
}

func (r *userRepository) GetById(ctx context.Context, userId int64) (user.UserModel, error) {
	sql, _, err := database.QueryBuilder.
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
		return user.UserModel{}, errors.Wrap(errors.DatabaseError, err, "syntax error")
	}

	row := r.Conn(ctx).QueryRow(ctx, sql)

	model := user.UserModel{Id: userId}

	err = row.Scan(
		&model.FirstName,
		&model.LastName,
		&model.Email,
		&model.Password,
	)
	if err != nil {
		return user.UserModel{}, parseGetUserByIdError(userId, err)
	}

	return model, nil
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (user.UserModel, error) {
	sql, _, err := database.QueryBuilder.
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
		return user.UserModel{}, errors.Wrap(errors.DatabaseError, err, "syntax error")
	}

	row := r.Conn(ctx).QueryRow(ctx, sql)

	model := user.UserModel{Email: email}

	err = row.Scan(
		&model.Id,
		&model.FirstName,
		&model.LastName,
		&model.Password,
	)
	if err != nil {
		return user.UserModel{}, parseGetUserByEmailError(email, err)
	}

	return model, nil
}
