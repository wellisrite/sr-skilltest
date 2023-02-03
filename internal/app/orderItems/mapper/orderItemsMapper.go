package mapper

import (
	"sr-skilltest/internal/app/orderItems"
	"sr-skilltest/internal/model/database"
	"sr-skilltest/internal/model/dto"
)

type OrderItemsMapper struct{}

func NewOrderItemsMapper() orderItems.OrderItemsMapper {
	return &OrderItemsMapper{}
}

func (m *OrderItemsMapper) ToResponseListPagination(orderItems *[]database.OrderItems, page int, pageLimit int, totalCount int) *dto.ResponsePagination {
	return &dto.ResponsePagination{
		Data:       orderItems,
		TotalCount: totalCount,
		Page:       page,
		PageLimit:  pageLimit,
	}
}

func (m *OrderItemsMapper) ToCreateOrderItems(payload *dto.RequestCreateOrderItems) (orderItems *database.OrderItems) {
	return &database.OrderItems{
		Name:  payload.Name,
		Price: payload.Price,
	}
}

func (m *OrderItemsMapper) ToUpdateOrderItems(payload *dto.RequestUpdateOrderItems) (orderItems *database.OrderItems) {
	return &database.OrderItems{
		Name:  payload.Name,
		Price: payload.Price,
	}
}
