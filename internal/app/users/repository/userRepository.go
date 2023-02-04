package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"sr-skilltest/internal/app/users"
	"sr-skilltest/internal/model/constant"
	"sr-skilltest/internal/model/database"

	"github.com/go-redis/redis"

	"gorm.io/gorm"
)

const CLASS = "user"

type UserRepository struct {
	DB    *gorm.DB
	Cache *redis.Client
}

func NewUserRepository(db *gorm.DB, cache *redis.Client) users.UserRepository {
	return &UserRepository{DB: db, Cache: cache}
}

func (r *UserRepository) GetByID(id uint64) (*database.User, error) {
	// Try to get data from cache
	var user database.User
	key := fmt.Sprintf("%s:%d", CLASS, id)
	val, err := r.Cache.Get(key).Bytes()
	if err == nil {
		if err := json.Unmarshal(val, &user); err == nil {
			return &user, nil
		}
	}

	result := r.DB.First(&user, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, result.Error
	}

	val, err = json.Marshal(user)
	if err != nil {
		return nil, err
	}
	if err := r.Cache.Set(key, val, 0).Err(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetAll(offset int, limit int) ([]database.User, int64, error) {
	// Try to get data from cache
	var totalCount int64
	var users []database.User

	cacheKey := fmt.Sprintf("users:%d:%d", offset, limit)
	cachedUsers, err := r.Cache.Get(cacheKey).Result()
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

	if err := r.Cache.Set(cacheKey, cached, constant.PAGINATION_CACHE_EXP_TIME.Abs()).Err(); err != nil {
		return users, totalCount, err
	}

	return users, totalCount, nil
}

func (r *UserRepository) Create(user *database.User) error {
	result := r.DB.Create(user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *UserRepository) Update(user *database.User, id uint64) error {
	result := r.DB.Where("id = ?", id).Updates(user)
	if result.Error != nil {
		return result.Error
	}

	key := fmt.Sprintf("%s:%d", CLASS, id)
	return r.Cache.Del(key).Err()
}

func (r *UserRepository) Delete(id uint64) error {
	result := r.DB.Where("id = ?", id).Delete(&database.User{})
	if result.Error != nil {
		return result.Error
	}

	key := fmt.Sprintf("%s:%d", CLASS, id)
	return r.Cache.Del(key).Err()
}
