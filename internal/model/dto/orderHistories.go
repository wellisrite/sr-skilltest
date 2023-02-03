package dto

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
