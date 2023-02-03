package dto

type RequestCreateUser struct {
	Name string `json:"name"`
}

type RequestUpdateUser struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
