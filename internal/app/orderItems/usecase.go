package orderItems

import (
	"github.com/labstack/echo"
)

// UserUsecase defines the behavior for managing users
type OrderItemsUsecase interface {
	Detail(traceID string, c echo.Context, id uint64) error
	List(traceID string, c echo.Context) error
	Create(traceID string, c echo.Context) error
	Update(traceID string, c echo.Context, id uint64) error
	Delete(traceID string, c echo.Context, id uint64) error
}
