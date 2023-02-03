package database

import (
	"time"

	"gorm.io/gorm"
)

type OrderItems struct {
	gorm.Model
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	ExpiredAt time.Time `json:"expired_at"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`
}
