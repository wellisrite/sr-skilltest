package users

import (
	"sr-skilltest/internal/domain"
	"sr-skilltest/internal/domain/database"
)

type UserRepository interface {
	GetByID(id uint64) (*domain.User, error)
	GetAll(offset int, limit int) ([]database.User, int64, error)
	Create(user *domain.User) error
	Update(user *domain.User, id uint64) error
	Delete(id uint64) error
}
