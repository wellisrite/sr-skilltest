package domain

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FullName   string    `json:"full_name"`
	FirstOrder time.Time `json:"first_order,omitempty"`
}
