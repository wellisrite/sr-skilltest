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

const CLASS = "orderItems"
const CACHETOTALCOUNTKEY = "orderItems:total"

type OrderItemsRepository struct {
	DB    *gorm.DB
	Cache *redis.Client
}

func NewOrderItemsRepository(db *gorm.DB, cache *redis.Client) domain.OrderItemsRepository {
	return &OrderItemsRepository{DB: db, Cache: cache}
}

func (r *OrderItemsRepository) GetByID(id uint64) (*domain.OrderItems, error) {
	var orderItems domain.OrderItems
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
	if err := r.Cache.Set(key, val, constant.ENTITY_CACHE_EXP_TIME).Err(); err != nil {
		return nil, err
	}

	return &orderItems, nil
}

func (r *OrderItemsRepository) GetAll(offset int, limit int) ([]domain.OrderItems, int64, error) {
	var orderItems []domain.OrderItems
	var totalCount int64

	temp, err := r.Cache.Get(CACHETOTALCOUNTKEY).Result()
	if err != nil && err == redis.Nil {
		r.DB.Model(&domain.OrderItems{}).Count(&totalCount)
	} else if err != nil {
		return nil, totalCount, err
	}

	if err == nil {
		totalCount, err = strconv.ParseInt(temp, 10, 64)
		if err != nil {
			return nil, totalCount, err
		}
	}

	result := r.DB.Limit(limit).Offset(offset).Find(&orderItems)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	r.DB.Model(&domain.OrderItems{}).Count(&totalCount)

	if err := r.Cache.Set(CACHETOTALCOUNTKEY, totalCount, constant.ENTITY_CACHE_EXP_TIME).Err(); err != nil {
		return nil, totalCount, err
	}

	return orderItems, totalCount, nil
}

func (r *OrderItemsRepository) Create(orderItems *domain.OrderItems) error {
	result := r.DB.Create(orderItems)
	if result.Error != nil {
		return result.Error
	}

	key := fmt.Sprintf("%s:%d", CLASS, orderItems.ID)
	val, err := json.Marshal(orderItems)
	if err != nil {
		return err
	}

	if err := r.Cache.Del(CACHETOTALCOUNTKEY).Err(); err != nil {
		return err
	}

	return r.Cache.Set(key, val, constant.ENTITY_CACHE_EXP_TIME).Err()
}

func (r *OrderItemsRepository) Update(orderItems *domain.OrderItems, id uint64) error {
	result := r.DB.
		Model(orderItems).
		Clauses(clause.Returning{}).
		Where("id = ?", id).
		Updates(orderItems)

	if result.Error != nil {
		return result.Error
	}

	key := fmt.Sprintf("%s:%d", CLASS, orderItems.ID)
	val, err := json.Marshal(orderItems)
	if err != nil {
		return err
	}

	if err := r.Cache.Del(CACHETOTALCOUNTKEY).Err(); err != nil {
		return err
	}

	return r.Cache.Set(key, val, constant.ENTITY_CACHE_EXP_TIME).Err()
}

func (r *OrderItemsRepository) Delete(id uint64) error {
	result := r.DB.Where("id = ?", id).Delete(&domain.OrderItems{})
	if result.Error != nil {
		return result.Error
	}

	if err := r.Cache.Del(CACHETOTALCOUNTKEY).Err(); err != nil {
		return err
	}

	key := fmt.Sprintf("%s:%d", CLASS, id)
	return r.Cache.Del(key).Err()
}
