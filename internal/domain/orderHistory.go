package domain

import (
	"sr-skilltest/internal/domain/dto"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type OrderHistories struct {
	gorm.Model
	UserID       uint   `json:"user_id"`
	OrderItemID  uint   `json:"order_item_id"`
	Descriptions string `json:"descriptions"`
	User         *User
	OrderItem    *OrderItems
}

// OrderHistoriesUsecase defines the behavior for managing users
type OrderHistoriesUsecase interface {
	Detail(traceID string, c echo.Context, id uint64) error
	List(traceID string, c echo.Context) error
	Create(traceID string, c echo.Context) error
	Update(traceID string, c echo.Context, id uint64) error
	Delete(traceID string, c echo.Context, id uint64) error
}

type OrderHistoriesRepository interface {
	GetByID(id uint64) (*OrderHistories, error)
	GetAll(offset int, limit int) ([]OrderHistories, int64, error)
	Create(traceID string, OrderHistories *OrderHistories) error
	Update(OrderHistories *OrderHistories, id uint64) error
	Delete(id uint64) error
}

type OrderHistoriesMapper interface {
	ToResponseListPagination(orderHistories *[]OrderHistories, page int, pageLimit int, totalCount int) *dto.ResponsePagination
	ToResponseGetByID(orderHistories *OrderHistories) *dto.ResponseGetOrderHistories
	ToCreateOrderHistories(payload *dto.RequestCreateOrderHistories) (OrderHistories *OrderHistories)
	ToUpdateOrderHistories(payload *dto.RequestUpdateOrderHistories) (OrderHistories *OrderHistories)
}
