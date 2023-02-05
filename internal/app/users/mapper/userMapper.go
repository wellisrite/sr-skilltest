package mapper

import (
	"sr-skilltest/internal/app/users"
	"sr-skilltest/internal/domain"
	"sr-skilltest/internal/domain/dto"
)

type UserMapper struct{}

func NewUserMapper() users.UserMapper {
	return &UserMapper{}
}

func (m *UserMapper) ToResponseListPagination(users *[]domain.User, page int, pageLimit int, totalCount int) *dto.ResponsePagination {
	var response []*dto.ResponseGetUser
	for _, user := range *users {
		response = append(response, &dto.ResponseGetUser{
			ID:         user.ID,
			CreatedAt:  user.CreatedAt,
			DeletedAt:  user.DeletedAt.Time,
			UpdatedAt:  user.UpdatedAt,
			FirstOrder: user.FirstOrder,
			FullName:   user.FullName,
		})
	}

	return &dto.ResponsePagination{
		Data:       response,
		TotalCount: totalCount,
		Page:       page,
		PageLimit:  pageLimit,
	}
}

func (m *UserMapper) ToResponseGetByID(user *domain.User) *dto.ResponseGetUser {
	return &dto.ResponseGetUser{
		ID:         user.ID,
		CreatedAt:  user.CreatedAt,
		DeletedAt:  user.DeletedAt.Time,
		UpdatedAt:  user.UpdatedAt,
		FirstOrder: user.FirstOrder,
		FullName:   user.FullName,
	}
}

func (m *UserMapper) ToCreateUser(payload *dto.RequestCreateUser) (user *domain.User) {
	return &domain.User{
		FullName: payload.Name,
	}
}

func (m *UserMapper) ToUpdateUser(payload *dto.RequestUpdateUser) (user *domain.User) {
	return &domain.User{
		FullName:   payload.Name,
		FirstOrder: payload.FirstOrder,
	}
}
