package dto

import "time"

type RequestCreateOrderHistories struct {
	UserID       int    `json:"user_id"`
	OrderItemID  int    `json:"order_item_id"`
	Descriptions string `json:"descriptions"`
}

type RequestUpdateOrderHistories struct {
	UserID       int    `json:"user_id"`
	OrderItemID  int    `json:"order_item_id"`
	Descriptions string `json:"descriptions"`
}

type ResponseGetOrderHistories struct {
	ID           uint                   `json:"id"`
	CreatedAt    time.Time              `json:"created_at"`
	UpdatedAt    time.Time              `json:"updated_at"`
	DeletedAt    time.Time              `json:"deleted_at"`
	Descriptions string                 `json:"description"`
	User         *ResponseGetUser       `json:"user"`
	OrderItem    *ResponseGetOrderItems `json:"order_item"`
}
