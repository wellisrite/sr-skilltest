package main

import (
	"sr-skilltest/internal/infra/log"

	_userHandler "sr-skilltest/internal/app/users/handler"
	_userMapper "sr-skilltest/internal/app/users/mapper"
	_userRepository "sr-skilltest/internal/app/users/repository"
	_userUsecase "sr-skilltest/internal/app/users/usecase"

	"github.com/labstack/echo"
)

func RunApplication() {
	properties := getProperties()

	log.SetupLogging(properties.App.Mode)

	e := echo.New()

	database := databaseConnect(properties)

	// Initialize dependencies, such as database and redis connections
	// db := repository.NewDB()
	// defer db.Close()
	// redis := repository.NewRedis()
	// defer redis.Close()
	userRepository := _userRepository.NewUserRepository(database)
	userMapper := _userMapper.NewUserMapper()
	userUsecase := _userUsecase.NewUserUsecase(userRepository, userMapper)
	_userHandler.NewUserHandler(e, userUsecase)

	// Start the server
	e.Start(":8080")
}
