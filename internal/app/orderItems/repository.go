package orderItems

import (
	"sr-skilltest/internal/model/database"
)

type OrderItemsRepository interface {
	GetByID(id uint64) (*database.OrderItems, error)
	GetAll(offset int, limit int) ([]database.OrderItems, int64, error)
	Create(OrderItems *database.OrderItems) error
	Update(OrderItems *database.OrderItems, id uint64) error
	Delete(id uint64) error
}
