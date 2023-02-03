package repository

import (
	"encoding/json"
	"errors"
	"sr-skilltest/internal/app/orderHistories"
	"sr-skilltest/internal/model/database"
	"time"

	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

type OrderHistoriesRepository struct {
	DB    *gorm.DB
	Cache *redis.Client
}

func NewOrderHistoriesRepository(db *gorm.DB, cache *redis.Client) orderHistories.OrderHistoriesRepository {
	return &OrderHistoriesRepository{DB: db, Cache: cache}
}

func (r *OrderHistoriesRepository) GetByID(id uint64) (*database.OrderHistories, error) {
	var orderHistories database.OrderHistories
	result := r.DB.First(&orderHistories, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, result.Error
	}

	return &orderHistories, nil
}

func (r *OrderHistoriesRepository) GetAll(offset int, limit int) ([]database.OrderHistories, int64, error) {
	var orderHistories []database.OrderHistories
	var totalCount int64

	cachedOrderHistories, err := r.Cache.Get("orderHistories").Result()
	if err == nil {
		if err := json.Unmarshal([]byte(cachedOrderHistories), &orderHistories); err != nil {
			return nil, totalCount, err
		}
		return orderHistories, totalCount, nil
	}
	result := r.DB.Limit(limit).Offset(offset).Find(&orderHistories)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	r.DB.Model(&database.OrderHistories{}).Count(&totalCount)

	// Save data to cache
	cached, err := json.Marshal(orderHistories)
	if err != nil {
		return nil, totalCount, err
	}
	if err := r.Cache.Set("orderHistories", cached, 600000).Err(); err != nil {
		return nil, totalCount, err
	}

	return orderHistories, totalCount, nil
}

func (r *OrderHistoriesRepository) Create(orderHistories *database.OrderHistories, user *database.User) error {
	tx := r.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Create(orderHistories).Error; err != nil {
		tx.Rollback()
		return err
	}

	if user.FirstOrder.IsZero() {
		if err := tx.Model(user).Where("id = ?", user.ID).Update("first_order", time.Now()).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	// Clear cache
	if err := r.Cache.Del("orderHistories").Err(); err != nil {
		return err
	}

	return tx.Commit().Error
}

func (r *OrderHistoriesRepository) Update(orderHistories *database.OrderHistories, id uint64) error {
	result := r.DB.Where("id = ?", id).Updates(orderHistories)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *OrderHistoriesRepository) Delete(id uint64) error {
	result := r.DB.Where("id = ?", id).Delete(&database.OrderHistories{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
