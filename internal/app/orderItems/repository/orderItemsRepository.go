package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"sr-skilltest/internal/app/orderItems"
	"sr-skilltest/internal/model/constant"
	"sr-skilltest/internal/model/database"

	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

const CLASS = "orderItems"

type OrderItemsRepository struct {
	DB    *gorm.DB
	Cache *redis.Client
}

func NewOrderItemsRepository(db *gorm.DB, cache *redis.Client) orderItems.OrderItemsRepository {
	return &OrderItemsRepository{DB: db, Cache: cache}
}

func (r *OrderItemsRepository) GetByID(id uint64) (*database.OrderItems, error) {
	var orderItems database.OrderItems
	// Try to get data from cache
	key := fmt.Sprintf("%s:%d", CLASS, id)
	val, err := r.Cache.Get(key).Bytes()
	if err == nil {
		if err := json.Unmarshal(val, &orderItems); err == nil {
			return &orderItems, nil
		}
	}

	result := r.DB.First(&orderItems, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, result.Error
	}

	val, err = json.Marshal(orderItems)
	if err != nil {
		return nil, err
	}
	if err := r.Cache.Set(key, val, 0).Err(); err != nil {
		return nil, err
	}

	return &orderItems, nil
}

func (r *OrderItemsRepository) GetAll(offset int, limit int) ([]database.OrderItems, int64, error) {
	var orderItems []database.OrderItems
	var totalCount int64

	// Try to get data from cache
	cacheKey := fmt.Sprintf("orderItems:%d:%d", offset, limit)
	cachedOrderItems, err := r.Cache.Get(cacheKey).Result()
	if err == nil {
		if err := json.Unmarshal([]byte(cachedOrderItems), &orderItems); err != nil {
			return nil, totalCount, err
		}
		return orderItems, totalCount, nil
	}

	result := r.DB.Limit(limit).Offset(offset).Find(&orderItems)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	r.DB.Model(&database.OrderItems{}).Count(&totalCount)

	// Save data to cache
	cached, err := json.Marshal(orderItems)
	if err != nil {
		return nil, totalCount, err
	}
	r.Cache.Set(cacheKey, cached, constant.PAGINATION_CACHE_EXP_TIME.Abs())

	return orderItems, totalCount, nil
}

func (r *OrderItemsRepository) Create(orderItems *database.OrderItems) error {
	result := r.DB.Create(orderItems)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *OrderItemsRepository) Update(orderItems *database.OrderItems, id uint64) error {
	result := r.DB.Where("id = ?", id).Updates(orderItems)
	if result.Error != nil {
		return result.Error
	}

	key := fmt.Sprintf("%s:%d", CLASS, id)
	return r.Cache.Del(key).Err()
}

func (r *OrderItemsRepository) Delete(id uint64) error {
	result := r.DB.Where("id = ?", id).Delete(&database.OrderItems{})
	if result.Error != nil {
		return result.Error
	}

	key := fmt.Sprintf("%s:%d", CLASS, id)
	return r.Cache.Del(key).Err()
}
