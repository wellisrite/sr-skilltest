package database

import (
	"time"

	"gorm.io/gorm"
)

type OrderHistories struct {
	gorm.Model
	UserID       uint      `json:"user_id"`
	OrderItemID  uint      `json:"order_item_id"`
	Descriptions string    `json:"descriptions"`
	DeletedAt    time.Time `json:"deleted_at,omitempty"`
	User         *User
	OrderItem    *OrderItems
}
