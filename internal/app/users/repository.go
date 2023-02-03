package users

import (
	"sr-skilltest/internal/model/database"
)

type UserRepository interface {
	GetByID(id uint64) (*database.User, error)
	GetAll(offset int, limit int) ([]database.User, int64, error)
	Create(user *database.User) error
	Update(user *database.User) error
	Delete(id uint) error
}
