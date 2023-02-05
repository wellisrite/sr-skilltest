package users

import (
	"sr-skilltest/internal/domain"
	"sr-skilltest/internal/domain/dto"
)

type UserMapper interface {
	ToResponseListPagination(users *[]domain.User, page int, pageLimit int, totalCount int) *dto.ResponsePagination
	ToResponseGetByID(user *domain.User) *dto.ResponseGetUser
	ToCreateUser(payload *dto.RequestCreateUser) (user *domain.User)
	ToUpdateUser(payload *dto.RequestUpdateUser) (user *domain.User)
}
