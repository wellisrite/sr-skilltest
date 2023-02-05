package handler

import (
	"fmt"
	"sr-skilltest/internal/app/orderItems"
	"sr-skilltest/internal/domain/constant"
	middleware "sr-skilltest/internal/middlewares"
	"sr-skilltest/internal/utilities"
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

	orderItemsRoutes.Use(middleware.LoggerMiddleware)
	orderItemsRoutes.GET("", handler.List)
	orderItemsRoutes.GET("/:id", handler.Detail)
	orderItemsRoutes.DELETE("/:id", handler.Delete)
	orderItemsRoutes.POST("", handler.Create)
	orderItemsRoutes.PUT("/:id", handler.Update)
}

// List returns a list of all orderItems
func (h *orderItemsHandler) List(c echo.Context) error {
	traceId, _ := c.Get(constant.CONTEXT_LOCALS_KEY_TRACE_ID).(string)
	if traceId == "" {
		traceId = utilities.CreateTraceID()
	}

	return h.orderItemsUsecase.List(traceId, c)
}

// Detail returns the details of a specific OrderItems
func (h *orderItemsHandler) Detail(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return fmt.Errorf("Not supported param")
	}

	traceId, _ := c.Get(constant.CONTEXT_LOCALS_KEY_TRACE_ID).(string)
	if traceId == "" {
		traceId = utilities.CreateTraceID()
	}

	return h.orderItemsUsecase.Detail(traceId, c, id)
}

// Update updates the details of a specific OrderItems
func (h *orderItemsHandler) Update(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return fmt.Errorf("Not supported param")
	}

	traceId, _ := c.Get(constant.CONTEXT_LOCALS_KEY_TRACE_ID).(string)
	if traceId == "" {
		traceId = utilities.CreateTraceID()
	}

	return h.orderItemsUsecase.Update(traceId, c, id)
}

// Delete soft-deletes a specific OrderItems
func (h *orderItemsHandler) Delete(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return fmt.Errorf("Not supported param")
	}

	traceId, _ := c.Get(constant.CONTEXT_LOCALS_KEY_TRACE_ID).(string)
	if traceId == "" {
		traceId = utilities.CreateTraceID()
	}

	return h.orderItemsUsecase.Delete(traceId, c, id)
}

// Create creates a new OrderItems
func (h *orderItemsHandler) Create(c echo.Context) error {
	traceId, _ := c.Get(constant.CONTEXT_LOCALS_KEY_TRACE_ID).(string)
	if traceId == "" {
		traceId = utilities.CreateTraceID()
	}

	return h.orderItemsUsecase.Create(traceId, c)
}
