package orderItems

import (
	"sr-skilltest/internal/domain"
	"sr-skilltest/internal/domain/dto"
)

type OrderItemsMapper interface {
	ToResponseListPagination(orderItems *[]domain.OrderItems, page int, pageLimit int, totalCount int) *dto.ResponsePagination
	ToResponseGetByID(orderItems *domain.OrderItems) *dto.ResponseGetOrderItems
	ToCreateOrderItems(payload *dto.RequestCreateOrderItems) (OrderItems *domain.OrderItems)
	ToUpdateOrderItems(payload *dto.RequestUpdateOrderItems) (OrderItems *domain.OrderItems)
}
