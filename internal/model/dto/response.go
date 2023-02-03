package dto

type ResponsePagination struct {
	Data       interface{} `json:"data"`
	TotalCount int         `json:"count"`
	Page       int         `json:"page"`
	PageLimit  int         `json:"page_limit"`
}

type ResponseWithMessage struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}
