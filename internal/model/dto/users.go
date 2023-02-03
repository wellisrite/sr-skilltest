package dto

import "time"

type RequestCreateUser struct {
	Name string `json:"name"`
}

type RequestUpdateUser struct {
	FirstOrder time.Time `json:"first_order"`
	Name       string    `json:"name"`
}
