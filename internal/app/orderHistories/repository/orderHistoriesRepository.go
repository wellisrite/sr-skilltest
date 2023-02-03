package repository

import (
	"errors"
	"sr-skilltest/internal/app/orderHistories"
	"sr-skilltest/internal/model/database"
	"time"

	"gorm.io/gorm"
)

type OrderHistoriesRepository struct {
	DB *gorm.DB
}

func NewOrderHistoriesRepository(db *gorm.DB) orderHistories.OrderHistoriesRepository {
	return &OrderHistoriesRepository{DB: db}
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
	var orderHistoriess []database.OrderHistories
	result := r.DB.Limit(limit).Offset(offset).Find(&orderHistoriess)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	var totalCount int64
	r.DB.Model(&database.OrderHistories{}).Count(&totalCount)

	return orderHistoriess, totalCount, nil
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
