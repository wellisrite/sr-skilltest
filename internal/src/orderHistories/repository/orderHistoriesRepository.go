package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"sr-skilltest/internal/domain"
	"sr-skilltest/internal/domain/constant"
	"sr-skilltest/internal/infra/cuslogger"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

const CLASS = "orderHistories"

type OrderHistoriesRepository struct {
	DB    *gorm.DB
	Cache *redis.Client
}

func NewOrderHistoriesRepository(db *gorm.DB, cache *redis.Client) domain.OrderHistoriesRepository {
	return &OrderHistoriesRepository{DB: db, Cache: cache}
}

func (r *OrderHistoriesRepository) GetByID(id uint64) (*domain.OrderHistories, error) {
	var orderHistories domain.OrderHistories
	key := fmt.Sprintf("%s:%d", CLASS, id)
	val, err := r.Cache.Get(key).Bytes()
	if err == nil {
		if err := json.Unmarshal(val, &orderHistories); err == nil {
			return &orderHistories, nil
		}
	}

	result := r.DB.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Unscoped()
	}).Preload("OrderItem", func(db *gorm.DB) *gorm.DB {
		return db.Unscoped()
	}).First(&orderHistories, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, result.Error
	}

	val, err = json.Marshal(orderHistories)
	if err != nil {
		return nil, err
	}
	if err := r.Cache.Set(key, val, constant.ENTITY_CACHE_EXP_TIME_1).Err(); err != nil {
		return nil, err
	}

	return &orderHistories, nil
}

func (r *OrderHistoriesRepository) GetAll(offset int, limit int) ([]domain.OrderHistories, int64, error) {
	var orderHistories []domain.OrderHistories
	var totalCount int64

	cacheKey := fmt.Sprintf("orderHistories:%d:%d", offset, limit)
	cacheTotalCountKey := fmt.Sprintf("orderHistories:%d:%d:total", offset, limit)

	cachedOrderHistories, err := r.Cache.Get(cacheKey).Result()
	if err == nil {
		temp, err := r.Cache.Get(cacheTotalCountKey).Result()
		if err != nil {
			return nil, totalCount, err
		}

		cachedTotalCount, err := strconv.ParseInt(temp, 10, 64)
		if err != nil {
			return nil, totalCount, err
		}

		if err := json.Unmarshal([]byte(cachedOrderHistories), &orderHistories); err != nil {
			return nil, cachedTotalCount, err
		}
		return orderHistories, cachedTotalCount, nil
	}

	result := r.DB.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Unscoped()
	}).Preload("OrderItem", func(db *gorm.DB) *gorm.DB {
		return db.Unscoped()
	}).Limit(limit).Offset(offset).Find(&orderHistories)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	r.DB.Model(&domain.OrderHistories{}).Count(&totalCount)

	// Save data to cache
	cached, err := json.Marshal(orderHistories)
	if err != nil {
		return nil, totalCount, err
	}
	if err := r.Cache.Set(cacheKey, cached, constant.PAGINATION_CACHE_EXP_TIME.Abs()).Err(); err != nil {
		return nil, totalCount, err
	}
	if err := r.Cache.Set(cacheTotalCountKey, totalCount, constant.PAGINATION_CACHE_EXP_TIME).Err(); err != nil {
		return nil, totalCount, err
	}
	return orderHistories, totalCount, nil
}

func (r *OrderHistoriesRepository) Create(traceID string, orderHistories *domain.OrderHistories) error {
	tx := r.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	var user domain.User
	if err := tx.First(&user, orderHistories.UserID).Error; err != nil {
		tx.Rollback()
		return err
	}

	var orderItem domain.OrderItems
	if err := tx.First(&orderItem, orderHistories.OrderItemID).Error; err != nil {
		tx.Rollback()
		return err
	}

	// not allowing to buy expired product
	if orderItem.ExpiredAt.Before(time.Now()) && !orderItem.ExpiredAt.IsZero() {
		return fmt.Errorf(constant.ERR_EXPIRED_PRODUCT)
	}

	if err := tx.Create(orderHistories).Error; err != nil {
		tx.Rollback()
		return err
	}

	if user.FirstOrder.IsZero() {
		cuslogger.Event(traceID, "customer first buy")
		if err := tx.Model(&user).Update("first_order", time.Now()).Error; err != nil {
			tx.Rollback()
			return err
		}

		r.Cache.Del(fmt.Sprintf("user:%d", orderHistories.UserID))
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	key := fmt.Sprintf("%s:%d", CLASS, orderHistories.ID)
	val, err := json.Marshal(orderHistories)
	if err != nil {
		return err
	}

	return r.Cache.Set(key, val, constant.ENTITY_CACHE_EXP_TIME).Err()
}

func (r *OrderHistoriesRepository) Update(orderHistories *domain.OrderHistories, id uint64) error {
	result := r.DB.
		Where("id = ?", id).
		Updates(orderHistories)
	if result.Error != nil {
		return result.Error
	}

	result = r.DB.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Unscoped()
	}).Preload("OrderItem", func(db *gorm.DB) *gorm.DB {
		return db.Unscoped()
	}).First(&orderHistories, id)
	if result.Error != nil {
		return result.Error
	}

	key := fmt.Sprintf("%s:%d", CLASS, orderHistories.ID)
	val, err := json.Marshal(orderHistories)
	if err != nil {
		return err
	}

	return r.Cache.Set(key, val, constant.ENTITY_CACHE_EXP_TIME).Err()
}

func (r *OrderHistoriesRepository) Delete(id uint64) error {
	result := r.DB.Where("id = ?", id).Delete(&domain.OrderHistories{})
	if result.Error != nil {
		return result.Error
	}

	key := fmt.Sprintf("%s:%d", CLASS, id)
	return r.Cache.Del(key).Err()
}
