package handler

import (
	"fmt"
	"sr-skilltest/internal/app/users"
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

	userRoutes.GET("", handler.List)
	userRoutes.GET("/:id", handler.Detail)
	userRoutes.POST("", handler.Create)
}

// List returns a list of all users
func (h *UserHandler) List(c echo.Context) error {
	return h.userUsecase.ListUsers(c)
}

// Detail returns the details of a specific user
func (h *UserHandler) Detail(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return fmt.Errorf("Not supported param")
	}

	return h.userUsecase.Detail(c, id)
}

// // Update updates the details of a specific user
// func (h *UserHandler) Update(c echo.Context) error {
// 	id := c.Param("id")
// 	user, err := h.userUsecase.Update(id, c.Request().Body)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, err.Error())
// 	}
// 	return c.JSON(http.StatusOK, user)
// }

// // Delete soft-deletes a specific user
// func (h *UserHandler) Delete(c echo.Context) error {
// 	id := c.Param("id")
// 	err := h.userUsecase.Delete(id)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, err.Error())
// 	}
// 	return c.JSON(http.StatusNoContent, nil)
// }

// Create creates a new user
func (h *UserHandler) Create(c echo.Context) error {
	return h.userUsecase.Create(c)
}
