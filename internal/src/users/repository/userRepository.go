package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"sr-skilltest/internal/domain"
	"sr-skilltest/internal/domain/constant"
	"strconv"

	"gorm.io/gorm/clause"

	"github.com/go-redis/redis"

	"gorm.io/gorm"
)

const CLASS = "user"

type UserRepository struct {
	DB    *gorm.DB
	Cache *redis.Client
}

func NewUserRepository(db *gorm.DB, cache *redis.Client) domain.UserRepository {
	return &UserRepository{DB: db, Cache: cache}
}

func (r *UserRepository) GetByID(id uint64) (*domain.User, error) {
	// Try to get data from cache
	var user domain.User
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
	if err := r.Cache.Set(key, val, constant.ENTITY_CACHE_EXP_TIME).Err(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetAll(offset int, limit int) ([]domain.User, int64, error) {
	// Try to get data from cache
	var totalCount int64
	var users []domain.User

	cacheKey := fmt.Sprintf("users:%d:%d", offset, limit)
	cacheTotalCountKey := fmt.Sprintf("users:%d:%d:total", offset, limit)

	cachedUsers, err := r.Cache.Get(cacheKey).Result()
	if err == nil {
		temp, err := r.Cache.Get(cacheTotalCountKey).Result()
		if err != nil {
			return nil, totalCount, err
		}

		cachedTotalCount, err := strconv.ParseInt(temp, 10, 64)
		if err != nil {
			return nil, totalCount, err
		}

		if err := json.Unmarshal([]byte(cachedUsers), &users); err != nil {
			return nil, cachedTotalCount, err
		}
		return users, cachedTotalCount, nil
	}

	result := r.DB.Limit(limit).Offset(offset).Find(&users)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	r.DB.Model(&domain.User{}).Count(&totalCount)

	// Save data to cache
	cached, err := json.Marshal(users)
	if err != nil {
		return nil, totalCount, err
	}

	if err := r.Cache.Set(cacheKey, cached, constant.PAGINATION_CACHE_EXP_TIME.Abs()).Err(); err != nil {
		return users, totalCount, err
	}

	if err := r.Cache.Set(cacheTotalCountKey, totalCount, constant.PAGINATION_CACHE_EXP_TIME).Err(); err != nil {
		return users, totalCount, err
	}

	return users, totalCount, nil
}

func (r *UserRepository) Create(user *domain.User) error {
	result := r.DB.Create(user)
	if result.Error != nil {
		return result.Error
	}

	key := fmt.Sprintf("%s:%d", CLASS, user.ID)
	val, err := json.Marshal(user)
	if err != nil {
		return err
	}

	return r.Cache.Set(key, val, constant.ENTITY_CACHE_EXP_TIME).Err()
}

func (r *UserRepository) Update(user *domain.User, id uint64) error {
	result := r.DB.
		Model(user).
		Clauses(clause.Returning{}).
		Where("id = ?", id).
		Updates(user)

	if result.Error != nil {
		return result.Error
	}

	key := fmt.Sprintf("%s:%d", CLASS, user.ID)
	val, err := json.Marshal(user)
	if err != nil {
		return err
	}

	return r.Cache.Set(key, val, constant.ENTITY_CACHE_EXP_TIME).Err()
}

func (r *UserRepository) Delete(id uint64) error {
	result := r.DB.Where("id = ?", id).Delete(&domain.User{})
	if result.Error != nil {
		return result.Error
	}

	key := fmt.Sprintf("%s:%d", CLASS, id)
	return r.Cache.Del(key).Err()
}
