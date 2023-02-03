package repository

import (
	"errors"
	"sr-skilltest/internal/app/users"
	"sr-skilltest/internal/model/database"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) users.UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) GetByID(id uint64) (*database.User, error) {
	var user database.User
	result := r.DB.First(&user, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, result.Error
	}

	return &user, nil
}

func (r *UserRepository) GetAll(offset int, limit int) ([]database.User, int64, error) {
	var users []database.User
	result := r.DB.Limit(limit).Offset(offset).Find(&users)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	var totalCount int64
	r.DB.Model(&database.User{}).Count(&totalCount)

	return users, totalCount, nil
}

func (r *UserRepository) Create(user *database.User) error {
	result := r.DB.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *UserRepository) Update(user *database.User) error {
	result := r.DB.Save(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *UserRepository) Delete(id uint) error {
	result := r.DB.Where("id = ?", id).Delete(&database.User{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
