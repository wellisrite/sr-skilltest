package database

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID         uint      `json:"id"`
	FullName   string    `json:"full_name"`
	FirstOrder time.Time `json:"first_order,omitempty"`
	DeletedAt  time.Time `json:"deleted_at,omitempty"`
}
