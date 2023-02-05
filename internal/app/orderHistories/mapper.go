package orderHistories

import (
	"sr-skilltest/internal/domain"
	"sr-skilltest/internal/domain/dto"
)

type OrderHistoriesMapper interface {
	ToResponseListPagination(orderHistories *[]domain.OrderHistories, page int, pageLimit int, totalCount int) *dto.ResponsePagination
	ToResponseGetByID(orderHistories *domain.OrderHistories) *dto.ResponseGetOrderHistories
	ToCreateOrderHistories(payload *dto.RequestCreateOrderHistories) (OrderHistories *domain.OrderHistories)
	ToUpdateOrderHistories(payload *dto.RequestUpdateOrderHistories) (OrderHistories *domain.OrderHistories)
}
