package users

import (
	"github.com/labstack/echo"
)

// UserUsecase defines the behavior for managing users
type UserUsecase interface {
	Detail(traceID string, c echo.Context, id uint64) error
	ListUsers(traceID string, c echo.Context) error
	Create(traceID string, c echo.Context) error
	Update(traceID string, c echo.Context, id uint64) error
	Delete(traceID string, c echo.Context, id uint64) error
}
