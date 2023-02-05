package mapper

import (
	"sr-skilltest/internal/app/orderItems"
	"sr-skilltest/internal/domain"
	"sr-skilltest/internal/domain/constant"
	"sr-skilltest/internal/domain/dto"
	"time"
)

type OrderItemsMapper struct{}

func NewOrderItemsMapper() orderItems.OrderItemsMapper {
	return &OrderItemsMapper{}
}

func (m *OrderItemsMapper) ToResponseListPagination(orderItems *[]domain.OrderItems, page int, pageLimit int, totalCount int) *dto.ResponsePagination {
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

func (m *OrderItemsMapper) ToResponseGetByID(orderItem *domain.OrderItems) *dto.ResponseGetOrderItems {
	return &dto.ResponseGetOrderItems{
		ID:        orderItem.ID,
		CreatedAt: orderItem.CreatedAt,
		DeletedAt: orderItem.DeletedAt.Time,
		UpdatedAt: orderItem.UpdatedAt,
		Name:      orderItem.Name,
		Price:     orderItem.Price,
		ExpiredAt: orderItem.ExpiredAt,
	}
}

func (m *OrderItemsMapper) ToCreateOrderItems(payload *dto.RequestCreateOrderItems) (orderItems *domain.OrderItems) {
	date, _ := time.Parse(constant.YYYY_MM_DD, payload.ExpiryDate)

	return &domain.OrderItems{
		Name:      payload.Name,
		Price:     payload.Price,
		ExpiredAt: date,
	}
}

func (m *OrderItemsMapper) ToUpdateOrderItems(payload *dto.RequestUpdateOrderItems) (orderItems *domain.OrderItems) {
	date, _ := time.Parse(constant.YYYY_MM_DD, payload.ExpiryDate)

	return &domain.OrderItems{
		Name:      payload.Name,
		Price:     payload.Price,
		ExpiredAt: date,
	}
}
