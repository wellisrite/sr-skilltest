package handler

import (
	"fmt"
	"sr-skilltest/internal/app/orderHistories"
	middleware "sr-skilltest/internal/middlewares"
	"sr-skilltest/internal/model/constant"
	"sr-skilltest/internal/utilities"
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

	orderHistoriesRoutes.Use(middleware.LoggerMiddleware)
	orderHistoriesRoutes.GET("", handler.List)
	orderHistoriesRoutes.GET("/:id", handler.Detail)
	orderHistoriesRoutes.DELETE("/:id", handler.Delete)
	orderHistoriesRoutes.POST("", handler.Create)
	orderHistoriesRoutes.PUT("/:id", handler.Update)
}

// List returns a list of all orderHistoriess
func (h *orderHistoriesHandler) List(c echo.Context) error {
	traceId, _ := c.Get(constant.CONTEXT_LOCALS_KEY_TRACE_ID).(string)
	if traceId == "" {
		traceId = utilities.CreateTraceID()
	}

	return h.orderHistoriesUsecase.List(traceId, c)
}

// Detail returns the details of a specific OrderHistories
func (h *orderHistoriesHandler) Detail(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return fmt.Errorf("Not supported param")
	}

	traceId, _ := c.Get(constant.CONTEXT_LOCALS_KEY_TRACE_ID).(string)
	if traceId == "" {
		traceId = utilities.CreateTraceID()
	}

	return h.orderHistoriesUsecase.Detail(traceId, c, id)
}

// Update updates the details of a specific OrderHistories
func (h *orderHistoriesHandler) Update(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return fmt.Errorf("Not supported param")
	}

	traceId, _ := c.Get(constant.CONTEXT_LOCALS_KEY_TRACE_ID).(string)
	if traceId == "" {
		traceId = utilities.CreateTraceID()
	}
	return h.orderHistoriesUsecase.Update(traceId, c, id)
}

// Delete soft-deletes a specific OrderHistories
func (h *orderHistoriesHandler) Delete(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return fmt.Errorf("Not supported param")
	}

	traceId, _ := c.Get(constant.CONTEXT_LOCALS_KEY_TRACE_ID).(string)
	if traceId == "" {
		traceId = utilities.CreateTraceID()
	}

	return h.orderHistoriesUsecase.Delete(traceId, c, id)
}

// Create creates a new OrderHistories
func (h *orderHistoriesHandler) Create(c echo.Context) error {
	traceId, _ := c.Get(constant.CONTEXT_LOCALS_KEY_TRACE_ID).(string)
	if traceId == "" {
		traceId = utilities.CreateTraceID()
	}

	return h.orderHistoriesUsecase.Create(traceId, c)
}
