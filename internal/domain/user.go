package domain

import (
	"time"

	"sr-skilltest/internal/domain/dto"

	"github.com/labstack/echo"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FullName   string    `json:"full_name"`
	FirstOrder time.Time `json:"first_order,omitempty"`
}

// UserUsecase defines the behavior for managing users
type UserUsecase interface {
	Detail(traceID string, c echo.Context, id uint64) error
	ListUsers(traceID string, c echo.Context) error
	Create(traceID string, c echo.Context) error
	Update(traceID string, c echo.Context, id uint64) error
	Delete(traceID string, c echo.Context, id uint64) error
}

type UserRepository interface {
	GetByID(id uint64) (*User, error)
	GetAll(offset int, limit int) ([]User, int64, error)
	Create(user *User) error
	Update(user *User, id uint64) error
	Delete(id uint64) error
}

type UserMapper interface {
	ToResponseListPagination(users *[]User, page int, pageLimit int, totalCount int) *dto.ResponsePagination
	ToResponseGetByID(user *User) *dto.ResponseGetUser
	ToCreateUser(payload *dto.RequestCreateUser) (user *User)
	ToUpdateUser(payload *dto.RequestUpdateUser) (user *User)
}
