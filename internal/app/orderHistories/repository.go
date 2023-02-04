package orderHistories

import (
	"sr-skilltest/internal/model/database"
)

type OrderHistoriesRepository interface {
	GetByID(id uint64) (*database.OrderHistories, error)
	GetAll(offset int, limit int) ([]database.OrderHistories, int64, error)
	Create(traceID string, OrderHistories *database.OrderHistories) error
	Update(OrderHistories *database.OrderHistories, id uint64) error
	Delete(id uint64) error
}
