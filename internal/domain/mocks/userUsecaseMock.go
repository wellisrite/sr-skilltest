package mocks

import (
	"github.com/labstack/echo"
	"github.com/stretchr/testify/mock"
)

type UserUsecaseMock struct {
	mock.Mock
}

func (u *UserUsecaseMock) Detail(traceID string, c echo.Context, id uint64) error {
	args := u.Called(traceID, c, id)
	return args.Error(1)
}

func (u *UserUsecaseMock) ListUsers(traceID string, c echo.Context) error {
	args := u.Called(traceID, c)
	return args.Error(1)
}

func (u *UserUsecaseMock) Create(traceID string, c echo.Context) error {
	args := u.Called(traceID, c)
	return args.Error(1)
}

func (u *UserUsecaseMock) Update(traceID string, c echo.Context, id uint64) error {
	args := u.Called(traceID, c, id)
	return args.Error(1)
}

func (u *UserUsecaseMock) Delete(traceID string, c echo.Context, id uint64) error {
	args := u.Called(traceID, c, id)
	return args.Error(1)
}
