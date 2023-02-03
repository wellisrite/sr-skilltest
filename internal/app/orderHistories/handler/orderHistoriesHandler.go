package handler

import (
	"fmt"
	"sr-skilltest/internal/app/orderHistories"
	"strconv"

	"github.com/labstack/echo"
)

// orderHistoriesHandler is the struct for the OrderHistories handlers
type orderHistoriesHandler struct {
	orderHistoriesUsecase orderHistories.OrderHistoriesUsecase
}

func NewOrderHistoriesHandler(c *echo.Echo, orderHistoriesUsecase orderHistories.OrderHistoriesUsecase) {
	handler := &orderHistoriesHandler{orderHistoriesUsecase: orderHistoriesUsecase}
	orderHistoriesRoutes := c.Group("/order-histories")

	orderHistoriesRoutes.GET("", handler.List)
	orderHistoriesRoutes.GET("/:id", handler.Detail)
	orderHistoriesRoutes.DELETE("/:id", handler.Delete)
	orderHistoriesRoutes.POST("", handler.Create)
	orderHistoriesRoutes.PUT("/:id", handler.Update)
}

// List returns a list of all orderHistoriess
func (h *orderHistoriesHandler) List(c echo.Context) error {
	return h.orderHistoriesUsecase.List(c)
}

// Detail returns the details of a specific OrderHistories
func (h *orderHistoriesHandler) Detail(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return fmt.Errorf("Not supported param")
	}

	return h.orderHistoriesUsecase.Detail(c, id)
}

// Update updates the details of a specific OrderHistories
func (h *orderHistoriesHandler) Update(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return fmt.Errorf("Not supported param")
	}

	return h.orderHistoriesUsecase.Update(c, id)
}

// Delete soft-deletes a specific OrderHistories
func (h *orderHistoriesHandler) Delete(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return fmt.Errorf("Not supported param")
	}

	return h.orderHistoriesUsecase.Delete(c, id)
}

// Create creates a new OrderHistories
func (h *orderHistoriesHandler) Create(c echo.Context) error {
	return h.orderHistoriesUsecase.Create(c)
}
