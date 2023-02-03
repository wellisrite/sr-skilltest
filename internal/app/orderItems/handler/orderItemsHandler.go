package handler

import (
	"fmt"
	"sr-skilltest/internal/app/orderItems"
	"strconv"

	"github.com/labstack/echo"
)

// orderItemsHandler is the struct for the OrderItems handlers
type orderItemsHandler struct {
	orderItemsUsecase orderItems.OrderItemsUsecase
}

func NewOrderItemsHandler(c *echo.Echo, orderItemsUsecase orderItems.OrderItemsUsecase) {
	handler := &orderItemsHandler{orderItemsUsecase: orderItemsUsecase}
	orderItemsRoutes := c.Group("/order-items")

	orderItemsRoutes.GET("", handler.List)
	orderItemsRoutes.GET("/:id", handler.Detail)
	orderItemsRoutes.DELETE("/:id", handler.Delete)
	orderItemsRoutes.POST("", handler.Create)
	orderItemsRoutes.PUT("/:id", handler.Update)
}

// List returns a list of all orderItemss
func (h *orderItemsHandler) List(c echo.Context) error {
	return h.orderItemsUsecase.List(c)
}

// Detail returns the details of a specific OrderItems
func (h *orderItemsHandler) Detail(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return fmt.Errorf("Not supported param")
	}

	return h.orderItemsUsecase.Detail(c, id)
}

// Update updates the details of a specific OrderItems
func (h *orderItemsHandler) Update(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return fmt.Errorf("Not supported param")
	}

	return h.orderItemsUsecase.Update(c, id)
}

// Delete soft-deletes a specific OrderItems
func (h *orderItemsHandler) Delete(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return fmt.Errorf("Not supported param")
	}

	return h.orderItemsUsecase.Delete(c, id)
}

// Create creates a new OrderItems
func (h *orderItemsHandler) Create(c echo.Context) error {
	return h.orderItemsUsecase.Create(c)
}
