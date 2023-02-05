package dto

import "time"

type RequestCreateUser struct {
	Name string `json:"name"`
}

type RequestUpdateUser struct {
	FirstOrder time.Time `json:"first_order"`
	Name       string    `json:"name"`
}

type ResponseGetUser struct {
	ID         uint      `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  time.Time `json:"deleted_at"`
	FirstOrder time.Time `json:"first_order"`
	FullName   string    `json:"full_name"`
}
