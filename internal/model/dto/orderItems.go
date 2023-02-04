package dto

import "time"

type RequestCreateOrderItems struct {
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	ExpiryDate string  `json:"expiry_date"`
}

type RequestUpdateOrderItems struct {
	Price      float64 `json:"price"`
	Name       string  `json:"name"`
	ExpiryDate string  `json:"expiry_date"`
}

type ResponseGetOrderItems struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	ExpiredAt time.Time `json:"expired_at"`
}
