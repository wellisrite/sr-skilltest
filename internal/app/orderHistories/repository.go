package orderHistories

import (
	"sr-skilltest/internal/domain"
	"sr-skilltest/internal/domain/database"
)

type OrderHistoriesRepository interface {
	GetByID(id uint64) (*domain.OrderHistories, error)
	GetAll(offset int, limit int) ([]database.OrderHistories, int64, error)
	Create(traceID string, OrderHistories *domain.OrderHistories) error
	Update(OrderHistories *domain.OrderHistories, id uint64) error
	Delete(id uint64) error
}
