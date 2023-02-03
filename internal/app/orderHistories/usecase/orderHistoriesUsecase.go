package usecases

import (
	"net/http"
	"sr-skilltest/internal/app/orderHistories"
	"sr-skilltest/internal/app/users"
	"sr-skilltest/internal/infra/cuslogger"
	"sr-skilltest/internal/model/dto"
	"strconv"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type OrderHistoriesUsecase struct {
	repo     orderHistories.OrderHistoriesRepository
	userRepo users.UserRepository
	mapper   orderHistories.OrderHistoriesMapper
}

// NewOrderHistoriesUsecase creates a new instance of OrderHistoriesUsecase
func NewOrderHistoriesUsecase(repo orderHistories.OrderHistoriesRepository, userRepo users.UserRepository, mapper orderHistories.OrderHistoriesMapper) orderHistories.OrderHistoriesUsecase {
	return &OrderHistoriesUsecase{
		repo:     repo,
		userRepo: userRepo,
		mapper:   mapper,
	}
}

// GetByID retrieves a orderHistories by its ID
func (u *OrderHistoriesUsecase) Detail(traceID string, c echo.Context, id uint64) error {
	orderHistories, err := u.repo.GetByID(id)
	if err != nil && err == gorm.ErrRecordNotFound {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	} else if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	cuslogger.Event("done", "processing request")
	return c.JSON(http.StatusOK, orderHistories)
}

func (u *OrderHistoriesUsecase) List(traceID string, c echo.Context) error {
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

	orderHistories, totalCount, err := u.repo.GetAll(offset, l)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve orderHistories"})
	}

	cuslogger.Event(traceID, "Done processing")
	return c.JSON(http.StatusOK, u.mapper.ToResponseListPagination(&orderHistories, p, l, int(totalCount)))
}

// Create creates a new orderHistories
func (u *OrderHistoriesUsecase) Create(traceID string, c echo.Context) error {
	request := &dto.RequestCreateOrderHistories{}
	if err := c.Bind(request); err != nil {
		return err
	}

	user, err := u.userRepo.GetByID(uint64(request.UserID))

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]string{"error": "Failed to get user"})
	}

	orderHistories := u.mapper.ToCreateOrderHistories(request)
	err = u.repo.Create(orderHistories, user)
	if err != nil {
		cuslogger.Error("create_order", err, "error in creating order")
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]string{"error": "Failed to create orderHistories"})
	}

	cuslogger.Event(traceID, "Done processing")
	return c.JSON(http.StatusCreated, dto.ResponseWithMessage{Status: true, Message: "OrderHistories has been created"})
}

// Update updates an existing orderHistories
func (u *OrderHistoriesUsecase) Update(traceID string, c echo.Context, id uint64) error {
	request := &dto.RequestUpdateOrderHistories{}
	if err := c.Bind(request); err != nil {
		return err
	}

	orderHistories := u.mapper.ToUpdateOrderHistories(request)
	err := u.repo.Update(orderHistories, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]string{"error": "Failed to update orderHistories"})
	}

	cuslogger.Event(traceID, "Done processing")
	return c.JSON(http.StatusOK, dto.ResponseWithMessage{Status: true, Message: "OrderHistories has been updated"})
}

// Delete deletes a orderHistories
func (u *OrderHistoriesUsecase) Delete(traceID string, c echo.Context, id uint64) error {
	err := u.repo.Delete(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]string{"error": "Failed to delete orderHistories"})
	}

	cuslogger.Event(traceID, "Done processing")
	return c.JSON(http.StatusOK, dto.ResponseWithMessage{Status: true, Message: "OrderHistories has been deleted"})
}
