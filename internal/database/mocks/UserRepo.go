// Code generated by mockery v2.10.4. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	model "go-backend-template/internal/model"
)

// UserRepo is an autogenerated mock type for the UserRepo type
type UserRepo struct {
	mock.Mock
}

// Add provides a mock function with given fields: ctx, user
func (_m *UserRepo) Add(ctx context.Context, user model.User) (int64, error) {
	ret := _m.Called(ctx, user)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, model.User) int64); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.User) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByEmail provides a mock function with given fields: ctx, email
func (_m *UserRepo) GetByEmail(ctx context.Context, email string) (model.User, error) {
	ret := _m.Called(ctx, email)

	var r0 model.User
	if rf, ok := ret.Get(0).(func(context.Context, string) model.User); ok {
		r0 = rf(ctx, email)
	} else {
		r0 = ret.Get(0).(model.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetById provides a mock function with given fields: ctx, userId
func (_m *UserRepo) GetById(ctx context.Context, userId int64) (model.User, error) {
	ret := _m.Called(ctx, userId)

	var r0 model.User
	if rf, ok := ret.Get(0).(func(context.Context, int64) model.User); ok {
		r0 = rf(ctx, userId)
	} else {
		r0 = ret.Get(0).(model.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, user
func (_m *UserRepo) Update(ctx context.Context, user model.User) (int64, error) {
	ret := _m.Called(ctx, user)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, model.User) int64); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.User) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
