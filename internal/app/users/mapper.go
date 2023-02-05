package users

import (
	"sr-skilltest/internal/model/database"
	"sr-skilltest/internal/model/dto"
)

type UserMapper interface {
	ToResponseListPagination(users *[]database.User, page int, pageLimit int, totalCount int) *dto.ResponsePagination
	ToResponseGetByID(user *database.User) *dto.ResponseGetUser
	ToCreateUser(payload *dto.RequestCreateUser) (user *database.User)
	ToUpdateUser(payload *dto.RequestUpdateUser) (user *database.User)
}
