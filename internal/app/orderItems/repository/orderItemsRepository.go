package repository

import (
	"errors"
	"sr-skilltest/internal/app/orderItems"
	"sr-skilltest/internal/model/database"

	"gorm.io/gorm"
)

type OrderItemsRepository struct {
	DB *gorm.DB
}

func NewOrderItemsRepository(db *gorm.DB) orderItems.OrderItemsRepository {
	return &OrderItemsRepository{DB: db}
}

func (r *OrderItemsRepository) GetByID(id uint64) (*database.OrderItems, error) {
	var orderItems database.OrderItems
	result := r.DB.First(&orderItems, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, result.Error
	}

	return &orderItems, nil
}

func (r *OrderItemsRepository) GetAll(offset int, limit int) ([]database.OrderItems, int64, error) {
	var orderItemss []database.OrderItems
	result := r.DB.Limit(limit).Offset(offset).Find(&orderItemss)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	var totalCount int64
	r.DB.Model(&database.OrderItems{}).Count(&totalCount)

	return orderItemss, totalCount, nil
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
	return nil
}

func (r *OrderItemsRepository) Delete(id uint64) error {
	result := r.DB.Where("id = ?", id).Delete(&database.OrderItems{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
