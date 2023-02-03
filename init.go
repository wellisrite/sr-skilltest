package main

import (
	"sr-skilltest/internal/infra/log"

	_userHandler "sr-skilltest/internal/app/users/handler"
	_userMapper "sr-skilltest/internal/app/users/mapper"
	_userRepository "sr-skilltest/internal/app/users/repository"
	_userUsecase "sr-skilltest/internal/app/users/usecase"

	_orderItemsHandler "sr-skilltest/internal/app/orderItems/handler"
	_orderItemsMapper "sr-skilltest/internal/app/orderItems/mapper"
	_orderItemsRepository "sr-skilltest/internal/app/orderItems/repository"
	_orderItemsUsecase "sr-skilltest/internal/app/orderItems/usecase"

	_orderHistoriesHandler "sr-skilltest/internal/app/orderHistories/handler"
	_orderHistoriesMapper "sr-skilltest/internal/app/orderHistories/mapper"
	_orderHistoriesRepository "sr-skilltest/internal/app/orderHistories/repository"
	_orderHistoriesUsecase "sr-skilltest/internal/app/orderHistories/usecase"

	"github.com/labstack/echo"
)

func RunApplication() {
	// Initialize dependencies, such as database and redis connections
	properties := getProperties()

	log.SetupLogging(properties.App.Mode)

	e := echo.New()
	database := databaseConnect(properties)

	// redis := repository.NewRedis()
	// defer redis.Close()

	userRepository := _userRepository.NewUserRepository(database)
	userMapper := _userMapper.NewUserMapper()
	userUsecase := _userUsecase.NewUserUsecase(userRepository, userMapper)
	_userHandler.NewUserHandler(e, userUsecase)

	orderItemsRepository := _orderItemsRepository.NewOrderItemsRepository(database)
	orderItemsMapper := _orderItemsMapper.NewOrderItemsMapper()
	orderItemsUsecase := _orderItemsUsecase.NewOrderItemsUsecase(orderItemsRepository, orderItemsMapper)
	_orderItemsHandler.NewOrderItemsHandler(e, orderItemsUsecase)

	orderHistoriesRepository := _orderHistoriesRepository.NewOrderHistoriesRepository(database)
	orderHistoriesMapper := _orderHistoriesMapper.NewOrderHistoriesMapper()
	orderHistoriesUsecase := _orderHistoriesUsecase.NewOrderHistoriesUsecase(orderHistoriesRepository, userRepository, orderHistoriesMapper)
	_orderHistoriesHandler.NewOrderHistoriesHandler(e, orderHistoriesUsecase)

	// Start the server
	e.Start(":8080")
}
