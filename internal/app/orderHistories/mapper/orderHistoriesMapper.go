package mapper

import (
	"sr-skilltest/internal/app/orderHistories"
	"sr-skilltest/internal/model/database"
	"sr-skilltest/internal/model/dto"
)

type OrderHistoriesMapper struct{}

func NewOrderHistoriesMapper() orderHistories.OrderHistoriesMapper {
	return &OrderHistoriesMapper{}
}

func (m *OrderHistoriesMapper) ToResponseListPagination(orderHistories *[]database.OrderHistories, page int, pageLimit int, totalCount int) *dto.ResponsePagination {
	return &dto.ResponsePagination{
		Data:       orderHistories,
		TotalCount: totalCount,
		Page:       page,
		PageLimit:  pageLimit,
	}
}

func (m *OrderHistoriesMapper) ToCreateOrderHistories(payload *dto.RequestCreateOrderHistories) (orderHistories *database.OrderHistories) {
	return &database.OrderHistories{
		UserID:       uint(payload.UserID),
		OrderItemID:  uint(payload.OrderItemID),
		Descriptions: payload.Descriptions,
	}
}

func (m *OrderHistoriesMapper) ToUpdateOrderHistories(payload *dto.RequestUpdateOrderHistories) (orderHistories *database.OrderHistories) {
	return &database.OrderHistories{
		UserID:       uint(payload.UserID),
		OrderItemID:  uint(payload.OrderItemID),
		Descriptions: payload.Descriptions,
	}
}
