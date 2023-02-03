package database

import (
	"gorm.io/gorm"
)

type OrderHistories struct {
	gorm.Model
	UserID       uint   `json:"user_id"`
	OrderItemID  uint   `json:"order_item_id"`
	Descriptions string `json:"descriptions"`
	User         *User
	OrderItem    *OrderItems
}
