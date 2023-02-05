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
	var response []*dto.ResponseGetOrderHistories
	for _, orderHistory := range *orderHistories {
		orderItem := orderHistory.OrderItem
		user := orderHistory.User
		var orderItemRes *dto.ResponseGetOrderItems
		var userResponse *dto.ResponseGetUser

		if orderItem != nil {
			orderItemRes = &dto.ResponseGetOrderItems{
				ID:        orderItem.ID,
				CreatedAt: orderItem.CreatedAt,
				DeletedAt: orderItem.DeletedAt.Time,
				UpdatedAt: orderItem.UpdatedAt,
				Name:      orderItem.Name,
				Price:     orderItem.Price,
				ExpiredAt: orderItem.ExpiredAt,
			}
		}

		if user != nil {
			userResponse = &dto.ResponseGetUser{
				ID:         user.ID,
				CreatedAt:  user.CreatedAt,
				DeletedAt:  user.DeletedAt.Time,
				UpdatedAt:  user.UpdatedAt,
				FirstOrder: user.FirstOrder,
				FullName:   user.FullName,
			}
		}

		response = append(response, &dto.ResponseGetOrderHistories{
			ID:           orderHistory.ID,
			CreatedAt:    orderHistory.CreatedAt,
			DeletedAt:    orderHistory.DeletedAt.Time,
			UpdatedAt:    orderHistory.UpdatedAt,
			Descriptions: orderHistory.Descriptions,
			User:         userResponse,
			OrderItem:    orderItemRes,
		})
	}

	return &dto.ResponsePagination{
		Data:       response,
		TotalCount: totalCount,
		Page:       page,
		PageLimit:  pageLimit,
	}
}

func (m *OrderHistoriesMapper) ToResponseGetByID(orderHistories *database.OrderHistories) *dto.ResponseGetOrderHistories {
	var response *dto.ResponseGetOrderHistories
	orderItem := orderHistories.OrderItem
	user := orderHistories.User
	var orderItemRes *dto.ResponseGetOrderItems
	var userResponse *dto.ResponseGetUser

	if orderItem != nil {
		orderItemRes = &dto.ResponseGetOrderItems{
			ID:        orderItem.ID,
			CreatedAt: orderItem.CreatedAt,
			DeletedAt: orderItem.DeletedAt.Time,
			UpdatedAt: orderItem.UpdatedAt,
			Name:      orderItem.Name,
			Price:     orderItem.Price,
			ExpiredAt: orderItem.ExpiredAt,
		}
	}

	if user != nil {
		userResponse = &dto.ResponseGetUser{
			ID:         user.ID,
			CreatedAt:  user.CreatedAt,
			DeletedAt:  user.DeletedAt.Time,
			UpdatedAt:  user.UpdatedAt,
			FirstOrder: user.FirstOrder,
			FullName:   user.FullName,
		}
	}

	response = &dto.ResponseGetOrderHistories{
		ID:           orderHistories.ID,
		CreatedAt:    orderHistories.CreatedAt,
		DeletedAt:    orderHistories.DeletedAt.Time,
		UpdatedAt:    orderHistories.UpdatedAt,
		Descriptions: orderHistories.Descriptions,
		User:         userResponse,
		OrderItem:    orderItemRes,
	}

	return response
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
