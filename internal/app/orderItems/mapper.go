package orderItems

import (
	"sr-skilltest/internal/model/database"
	"sr-skilltest/internal/model/dto"
)

type OrderItemsMapper interface {
	ToResponseListPagination(orderItems *[]database.OrderItems, page int, pageLimit int, totalCount int) *dto.ResponsePagination
	ToCreateOrderItems(payload *dto.RequestCreateOrderItems) (OrderItems *database.OrderItems)
	ToUpdateOrderItems(payload *dto.RequestUpdateOrderItems) (OrderItems *database.OrderItems)
}
