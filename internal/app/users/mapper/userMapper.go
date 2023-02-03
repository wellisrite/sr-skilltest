package mapper

import (
	"sr-skilltest/internal/app/users"
	"sr-skilltest/internal/model/database"
	"sr-skilltest/internal/model/dto"
)

type UserMapper struct{}

func NewUserMapper() users.UserMapper {
	return &UserMapper{}
}

func (m *UserMapper) ToResponseListPagination(users *[]database.User, page int, pageLimit int, totalCount int) *dto.ResponsePagination {
	return &dto.ResponsePagination{
		Data:       users,
		TotalCount: totalCount,
		Page:       page,
		PageLimit:  pageLimit,
	}
}

func (m *UserMapper) ToCreateUser(payload *dto.RequestCreateUser) (user *database.User) {
	return &database.User{
		FullName: payload.Name,
	}
}

func (m *UserMapper) ToUpdateUser(payload *dto.RequestUpdateUser) (user *database.User) {
	return &database.User{
		FullName:   payload.Name,
		FirstOrder: payload.FirstOrder,
	}
}
