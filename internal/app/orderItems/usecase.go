package orderItems

import (
	"github.com/labstack/echo"
)

// UserUsecase defines the behavior for managing users
type OrderItemsUsecase interface {
	Detail(c echo.Context, id uint64) error
	List(c echo.Context) error
	Create(c echo.Context) error
	Update(c echo.Context, id uint64) error
	Delete(c echo.Context, id uint64) error
}
