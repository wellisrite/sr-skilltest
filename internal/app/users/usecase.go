package users

import (
	"github.com/labstack/echo"
)

// UserUsecase defines the behavior for managing users
type UserUsecase interface {
	Detail(c echo.Context, id uint64) error
	ListUsers(c echo.Context) error
	Create(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context, id uint) error
}
