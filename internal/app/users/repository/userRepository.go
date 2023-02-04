package repository

import (
	"encoding/json"
	"errors"
	"sr-skilltest/internal/app/users"
	"sr-skilltest/internal/model/database"

	"github.com/go-redis/redis"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB    *gorm.DB
	Cache *redis.Client
}

func NewUserRepository(db *gorm.DB, cache *redis.Client) users.UserRepository {
	return &UserRepository{DB: db, Cache: cache}
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
	// Try to get data from cache
	var totalCount int64
	var users []database.User
	cachedUsers, err := r.Cache.Get("users").Result()
	if err == nil {
		if err := json.Unmarshal([]byte(cachedUsers), &users); err != nil {
			return nil, totalCount, err
		}
		return users, totalCount, nil
	}

	result := r.DB.Limit(limit).Offset(offset).Find(&users)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	r.DB.Model(&database.User{}).Count(&totalCount)

	// Save data to cache
	cached, err := json.Marshal(users)
	if err != nil {
		return nil, totalCount, err
	}

	r.Cache.Set("users", cached, 0)

	return users, totalCount, nil
}

func (r *UserRepository) Create(user *database.User) error {
	result := r.DB.Create(user)
	if result.Error != nil {
		return result.Error
	}

	// Clear cache
	r.Cache.Del("users")

	return nil
}

func (r *UserRepository) Update(user *database.User, id uint64) error {
	result := r.DB.Where("id = ?", id).Updates(user)
	if result.Error != nil {
		return result.Error
	}

	r.Cache.Del("users")

	return nil
}

func (r *UserRepository) Delete(id uint64) error {
	result := r.DB.Where("id = ?", id).Delete(&database.User{})
	if result.Error != nil {
		return result.Error
	}

	r.Cache.Del("users")

	return nil
}
