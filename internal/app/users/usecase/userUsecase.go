package usecases

import (
	"net/http"
	"sr-skilltest/internal/app/users"
	"sr-skilltest/internal/infra/cuslogger"
	"sr-skilltest/internal/model/dto"
	"strconv"
	"time"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type UserUsecase struct {
	repo   users.UserRepository
	mapper users.UserMapper
}

// NewUserUsecase creates a new instance of UserUsecase
func NewUserUsecase(repo users.UserRepository, mapper users.UserMapper) users.UserUsecase {
	return &UserUsecase{
		repo:   repo,
		mapper: mapper,
	}
}

// GetByID retrieves a user by its ID
func (u *UserUsecase) Detail(traceID string, c echo.Context, id uint64) error {
	user, err := u.repo.GetByID(id)
	if err != nil && err == gorm.ErrRecordNotFound {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	} else if err != nil {
		cuslogger.Error(traceID, err, err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	cuslogger.Event(time.Now().String(), traceID, " done processing\n")
	return c.JSON(http.StatusOK, u.mapper.ToResponseGetByID(user))
}

func (u *UserUsecase) ListUsers(traceID string, c echo.Context) error {
	page := c.QueryParam("page")
	if page == "" {
		page = "1"
	}

	limit := c.QueryParam("limit")
	if limit == "" {
		limit = "10"
	}

	p, err := strconv.Atoi(page)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]string{"error": "Invalid page parameter"})
	}

	l, err := strconv.Atoi(limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]string{"error": "Invalid limit parameter"})
	}

	offset := (p - 1) * l

	users, totalCount, err := u.repo.GetAll(offset, l)
	if err != nil {
		cuslogger.Error(traceID, err, err.Error())

		return echo.NewHTTPError(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve users"})
	}

	cuslogger.Event(time.Now().String(), traceID, " done processing\n")

	return c.JSON(http.StatusOK, u.mapper.ToResponseListPagination(&users, p, l, int(totalCount)))
}

// Create creates a new user
func (u *UserUsecase) Create(traceID string, c echo.Context) error {
	request := &dto.RequestCreateUser{}
	if err := c.Bind(request); err != nil {
		return err
	}

	user := u.mapper.ToCreateUser(request)
	err := u.repo.Create(user)
	if err != nil {
		cuslogger.Error(traceID, err, err.Error())

		return echo.NewHTTPError(http.StatusInternalServerError, map[string]string{"error": "Failed to create user"})
	}

	cuslogger.Event(time.Now().String(), traceID, " done processing\n")

	return c.JSON(http.StatusCreated, dto.ResponseWithMessage{Status: true, Message: "User has been created"})
}

// Update updates an existing user
func (u *UserUsecase) Update(traceID string, c echo.Context, id uint64) error {
	request := &dto.RequestUpdateUser{}
	if err := c.Bind(request); err != nil {
		return err
	}

	user := u.mapper.ToUpdateUser(request)
	err := u.repo.Update(user, id)
	if err != nil {
		cuslogger.Error(traceID, err, err.Error())

		return echo.NewHTTPError(http.StatusInternalServerError, map[string]string{"error": "Failed to update user"})
	}

	cuslogger.Event(time.Now().String(), traceID, " done processing\n")

	return c.JSON(http.StatusOK, dto.ResponseWithMessage{Status: true, Message: "User has been updated"})
}

// Delete deletes a user
func (u *UserUsecase) Delete(traceID string, c echo.Context, id uint64) error {
	err := u.repo.Delete(id)
	if err != nil {
		cuslogger.Error(traceID, err, err.Error())

		return echo.NewHTTPError(http.StatusInternalServerError, map[string]string{"error": "Failed to delete user"})
	}

	cuslogger.Event(time.Now().String(), traceID, " done processing\n")
	return c.JSON(http.StatusOK, dto.ResponseWithMessage{Status: true, Message: "User has been deleted"})
}
