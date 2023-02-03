package handler

import (
	"fmt"
	"sr-skilltest/internal/app/users"
	middleware "sr-skilltest/internal/middlewares"
	"sr-skilltest/internal/model/constant"
	"sr-skilltest/internal/utilities"
	"strconv"

	"github.com/labstack/echo"
)

// UserHandler is the struct for the user handlers
type UserHandler struct {
	userUsecase users.UserUsecase
}

func NewUserHandler(c *echo.Echo, userUsecase users.UserUsecase) {
	handler := &UserHandler{userUsecase: userUsecase}
	userRoutes := c.Group("/users")

	userRoutes.Use(middleware.LoggerMiddleware)

	userRoutes.GET("", handler.List)
	userRoutes.GET("/:id", handler.Detail)
	userRoutes.DELETE("/:id", handler.Delete)
	userRoutes.POST("", handler.Create)
	userRoutes.PUT("/:id", handler.Update)
}

// List returns a list of all users
func (h *UserHandler) List(c echo.Context) error {
	traceId, _ := c.Get(constant.CONTEXT_LOCALS_KEY_TRACE_ID).(string)
	if traceId == "" {
		traceId = utilities.CreateTraceID()
	}

	return h.userUsecase.ListUsers(traceId, c)
}

// Detail returns the details of a specific user
func (h *UserHandler) Detail(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return fmt.Errorf("Not supported param")
	}

	traceId, _ := c.Get(constant.CONTEXT_LOCALS_KEY_TRACE_ID).(string)
	if traceId == "" {
		traceId = utilities.CreateTraceID()
	}

	return h.userUsecase.Detail(traceId, c, id)
}

// Update updates the details of a specific user
func (h *UserHandler) Update(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return fmt.Errorf("Not supported param")
	}

	traceId, _ := c.Get(constant.CONTEXT_LOCALS_KEY_TRACE_ID).(string)
	if traceId == "" {
		traceId = utilities.CreateTraceID()
	}

	return h.userUsecase.Update(traceId, c, id)
}

// Delete soft-deletes a specific user
func (h *UserHandler) Delete(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return fmt.Errorf("Not supported param")
	}

	traceId, _ := c.Get(constant.CONTEXT_LOCALS_KEY_TRACE_ID).(string)
	if traceId == "" {
		traceId = utilities.CreateTraceID()
	}

	return h.userUsecase.Delete(traceId, c, id)
}

// Create creates a new user
func (h *UserHandler) Create(c echo.Context) error {
	traceId, _ := c.Get(constant.CONTEXT_LOCALS_KEY_TRACE_ID).(string)
	if traceId == "" {
		traceId = utilities.CreateTraceID()
	}

	return h.userUsecase.Create(traceId, c)
}
