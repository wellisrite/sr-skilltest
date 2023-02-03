package usecases

import (
	"net/http"
	"sr-skilltest/internal/app/orderItems"
	"sr-skilltest/internal/infra/cuslogger"
	"sr-skilltest/internal/model/dto"
	"strconv"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type OrderItemsUsecase struct {
	repo   orderItems.OrderItemsRepository
	mapper orderItems.OrderItemsMapper
}

// NewOrderItemsUsecase creates a new instance of OrderItemsUsecase
func NewOrderItemsUsecase(repo orderItems.OrderItemsRepository, mapper orderItems.OrderItemsMapper) orderItems.OrderItemsUsecase {
	return &OrderItemsUsecase{
		repo:   repo,
		mapper: mapper,
	}
}

// GetByID retrieves a orderItems by its ID
func (u *OrderItemsUsecase) Detail(traceID string, c echo.Context, id uint64) error {
	orderItems, err := u.repo.GetByID(id)
	if err != nil && err == gorm.ErrRecordNotFound {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	} else if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	cuslogger.Event(traceID, "Done processing")
	return c.JSON(http.StatusOK, orderItems)
}

func (u *OrderItemsUsecase) List(traceID string, c echo.Context) error {
	page := c.QueryParam("page")
	if page == "" {
		page = "1"
	}

	limit := c.QueryParam("limit")
	if limit == "" {
		limit = "10"
	}

	p, err := strconv.Atoi(page)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]string{"error": "Invalid page parameter"})
	}

	l, err := strconv.Atoi(limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]string{"error": "Invalid limit parameter"})
	}

	offset := (p - 1) * l

	orderItems, totalCount, err := u.repo.GetAll(offset, l)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve orderItems"})
	}

	cuslogger.Event(traceID, "Done processing")
	return c.JSON(http.StatusOK, u.mapper.ToResponseListPagination(&orderItems, p, l, int(totalCount)))
}

// Create creates a new orderItems
func (u *OrderItemsUsecase) Create(traceID string, c echo.Context) error {
	request := &dto.RequestCreateOrderItems{}
	if err := c.Bind(request); err != nil {
		return err
	}

	orderItems := u.mapper.ToCreateOrderItems(request)
	err := u.repo.Create(orderItems)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]string{"error": "Failed to create orderItems"})
	}

	cuslogger.Event(traceID, "Done processing")
	return c.JSON(http.StatusCreated, dto.ResponseWithMessage{Status: true, Message: "OrderItems has been created"})
}

// Update updates an existing orderItems
func (u *OrderItemsUsecase) Update(traceID string, c echo.Context, id uint64) error {
	request := &dto.RequestUpdateOrderItems{}
	if err := c.Bind(request); err != nil {
		return err
	}

	orderItems := u.mapper.ToUpdateOrderItems(request)
	err := u.repo.Update(orderItems, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]string{"error": "Failed to update orderItems"})
	}

	cuslogger.Event(traceID, "Done processing")
	return c.JSON(http.StatusOK, dto.ResponseWithMessage{Status: true, Message: "OrderItems has been updated"})
}

// Delete deletes a orderItems
func (u *OrderItemsUsecase) Delete(traceID string, c echo.Context, id uint64) error {
	err := u.repo.Delete(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]string{"error": "Failed to delete orderItems"})
	}

	cuslogger.Event(traceID, "Done processing")
	return c.JSON(http.StatusOK, dto.ResponseWithMessage{Status: true, Message: "OrderItems has been deleted"})
}
