package usecases

import (
	"net/http"
	"sr-skilltest/internal/domain"
	"sr-skilltest/internal/domain/dto"
	"sr-skilltest/internal/infra/cuslogger"
	"strconv"
	"time"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type OrderItemsUsecase struct {
	repo   domain.OrderItemsRepository
	mapper domain.OrderItemsMapper
}

// NewOrderItemsUsecase creates a new instance of OrderItemsUsecase
func NewOrderItemsUsecase(repo domain.OrderItemsRepository, mapper domain.OrderItemsMapper) domain.OrderItemsUsecase {
	return &OrderItemsUsecase{
		repo:   repo,
		mapper: mapper,
	}
}

// GetByID retrieves a domain by its ID
func (u *OrderItemsUsecase) Detail(traceID string, c echo.Context, id uint64) error {
	domain, err := u.repo.GetByID(id)
	if err != nil && err == gorm.ErrRecordNotFound {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	} else if err != nil {
		cuslogger.Error(traceID, err, err.Error())

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	cuslogger.Event(time.Now().String(), traceID, " done processing\n")
	return c.JSON(http.StatusOK, u.mapper.ToResponseGetByID(domain))
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

	domain, totalCount, err := u.repo.GetAll(offset, l)
	if err != nil {
		cuslogger.Error(traceID, err, err.Error())

		return echo.NewHTTPError(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve domain"})
	}

	cuslogger.Event(time.Now().String(), traceID, " done processing\n")
	return c.JSON(http.StatusOK, u.mapper.ToResponseListPagination(&domain, p, l, int(totalCount)))
}

// Create creates a new domain
func (u *OrderItemsUsecase) Create(traceID string, c echo.Context) error {
	request := &dto.RequestCreateOrderItems{}
	if err := c.Bind(request); err != nil {
		return err
	}

	domain := u.mapper.ToCreateOrderItems(request)
	err := u.repo.Create(domain)
	if err != nil {
		cuslogger.Error(traceID, err, err.Error())

		return echo.NewHTTPError(http.StatusInternalServerError, map[string]string{"error": "Failed to create domain"})
	}

	cuslogger.Event(time.Now().String(), traceID, " done processing\n")
	return c.JSON(http.StatusCreated, dto.ResponseWithMessage{Status: true, Message: "OrderItems has been created"})
}

// Update updates an existing domain
func (u *OrderItemsUsecase) Update(traceID string, c echo.Context, id uint64) error {
	request := &dto.RequestUpdateOrderItems{}
	if err := c.Bind(request); err != nil {
		return err
	}

	domain := u.mapper.ToUpdateOrderItems(request)
	err := u.repo.Update(domain, id)
	if err != nil {
		cuslogger.Error(traceID, err, err.Error())

		return echo.NewHTTPError(http.StatusInternalServerError, map[string]string{"error": "Failed to update domain"})
	}

	cuslogger.Event(time.Now().String(), traceID, " done processing\n")
	return c.JSON(http.StatusOK, dto.ResponseWithMessage{Status: true, Message: "OrderItems has been updated"})
}

// Delete deletes a domain
func (u *OrderItemsUsecase) Delete(traceID string, c echo.Context, id uint64) error {
	err := u.repo.Delete(id)
	if err != nil {
		cuslogger.Error(traceID, err, err.Error())

		return echo.NewHTTPError(http.StatusInternalServerError, map[string]string{"error": "Failed to delete domain"})
	}

	cuslogger.Event(time.Now().String(), traceID, " done processing\n")
	return c.JSON(http.StatusOK, dto.ResponseWithMessage{Status: true, Message: "OrderItems has been deleted"})
}
