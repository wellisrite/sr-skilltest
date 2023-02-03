package orderHistories

import (
	"sr-skilltest/internal/model/database"
	"sr-skilltest/internal/model/dto"
)

type OrderHistoriesMapper interface {
	ToResponseListPagination(orderHistories *[]database.OrderHistories, page int, pageLimit int, totalCount int) *dto.ResponsePagination
	ToCreateOrderHistories(payload *dto.RequestCreateOrderHistories) (OrderHistories *database.OrderHistories)
	ToUpdateOrderHistories(payload *dto.RequestUpdateOrderHistories) (OrderHistories *database.OrderHistories)
}
