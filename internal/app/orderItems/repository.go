package orderItems

import (
	"sr-skilltest/internal/domain"
)

type OrderItemsRepository interface {
	GetByID(id uint64) (*domain.OrderItems, error)
	GetAll(offset int, limit int) ([]domain.OrderItems, int64, error)
	Create(OrderItems *domain.OrderItems) error
	Update(OrderItems *domain.OrderItems, id uint64) error
	Delete(id uint64) error
}
