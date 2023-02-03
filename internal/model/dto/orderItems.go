package dto

type RequestCreateOrderItems struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type RequestUpdateOrderItems struct {
	Price float64 `json:"price"`
	Name  string  `json:"name"`
}
