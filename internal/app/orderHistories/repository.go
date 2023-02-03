package orderHistories

import (
	"sr-skilltest/internal/model/database"
)

type OrderHistoriesRepository interface {
	GetByID(id uint64) (*database.OrderHistories, error)
	GetAll(offset int, limit int) ([]database.OrderHistories, int64, error)
	Create(OrderHistories *database.OrderHistories, user *database.User) error
	Update(OrderHistories *database.OrderHistories, id uint64) error
	Delete(id uint64) error
}
