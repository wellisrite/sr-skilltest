package mapper

import (
	"sr-skilltest/internal/app/orderItems"
	"sr-skilltest/internal/model/constant"
	"sr-skilltest/internal/model/database"
	"sr-skilltest/internal/model/dto"
	"time"
)

type OrderItemsMapper struct{}

func NewOrderItemsMapper() orderItems.OrderItemsMapper {
	return &OrderItemsMapper{}
}

func (m *OrderItemsMapper) ToResponseListPagination(orderItems *[]database.OrderItems, page int, pageLimit int, totalCount int) *dto.ResponsePagination {
	var response []*dto.ResponseGetOrderItems
	for _, orderItem := range *orderItems {
		response = append(response, &dto.ResponseGetOrderItems{
			ID:        orderItem.ID,
			CreatedAt: orderItem.CreatedAt,
			DeletedAt: orderItem.DeletedAt.Time,
			UpdatedAt: orderItem.UpdatedAt,
			Name:      orderItem.Name,
			Price:     orderItem.Price,
			ExpiredAt: orderItem.ExpiredAt,
		})
	}

	return &dto.ResponsePagination{
		Data:       response,
		TotalCount: totalCount,
		Page:       page,
		PageLimit:  pageLimit,
	}
}

func (m *OrderItemsMapper) ToCreateOrderItems(payload *dto.RequestCreateOrderItems) (orderItems *database.OrderItems) {
	date, _ := time.Parse(constant.YYYY_MM_DD, payload.ExpiryDate)

	return &database.OrderItems{
		Name:      payload.Name,
		Price:     payload.Price,
		ExpiredAt: date,
	}
}

func (m *OrderItemsMapper) ToUpdateOrderItems(payload *dto.RequestUpdateOrderItems) (orderItems *database.OrderItems) {
	date, _ := time.Parse(constant.YYYY_MM_DD, payload.ExpiryDate)

	return &database.OrderItems{
		Name:      payload.Name,
		Price:     payload.Price,
		ExpiredAt: date,
	}
}
