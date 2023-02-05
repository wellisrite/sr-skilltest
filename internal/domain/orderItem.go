package domain

import (
	"sr-skilltest/internal/domain/dto"
	"time"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type OrderItems struct {
	gorm.Model
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	ExpiredAt time.Time `json:"expired_at"`
}

type OrderItemsUsecase interface {
	Detail(traceID string, c echo.Context, id uint64) error
	List(traceID string, c echo.Context) error
	Create(traceID string, c echo.Context) error
	Update(traceID string, c echo.Context, id uint64) error
	Delete(traceID string, c echo.Context, id uint64) error
}

type OrderItemsRepository interface {
	GetByID(id uint64) (*OrderItems, error)
	GetAll(offset int, limit int) ([]OrderItems, int64, error)
	Create(OrderItems *OrderItems) error
	Update(OrderItems *OrderItems, id uint64) error
	Delete(id uint64) error
}

type OrderItemsMapper interface {
	ToResponseListPagination(orderItems *[]OrderItems, page int, pageLimit int, totalCount int) *dto.ResponsePagination
	ToResponseGetByID(orderItems *OrderItems) *dto.ResponseGetOrderItems
	ToCreateOrderItems(payload *dto.RequestCreateOrderItems) (OrderItems *OrderItems)
	ToUpdateOrderItems(payload *dto.RequestUpdateOrderItems) (OrderItems *OrderItems)
}
