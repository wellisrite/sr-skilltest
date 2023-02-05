package main

import (
	"sr-skilltest/internal/infra/cuslogger"

	_userHandler "sr-skilltest/internal/src/users/delivery/http"
	_userMapper "sr-skilltest/internal/src/users/mapper"
	_userRepository "sr-skilltest/internal/src/users/repository"
	_userUsecase "sr-skilltest/internal/src/users/usecase"

	_orderItemsHandler "sr-skilltest/internal/src/orderItems/delivery/http"
	_orderItemsMapper "sr-skilltest/internal/src/orderItems/mapper"
	_orderItemsRepository "sr-skilltest/internal/src/orderItems/repository"
	_orderItemsUsecase "sr-skilltest/internal/src/orderItems/usecase"

	_orderHistoriesHandler "sr-skilltest/internal/src/orderHistories/delivery/http"
	_orderHistoriesMapper "sr-skilltest/internal/src/orderHistories/mapper"
	_orderHistoriesRepository "sr-skilltest/internal/src/orderHistories/repository"
	_orderHistoriesUsecase "sr-skilltest/internal/src/orderHistories/usecase"

	"github.com/labstack/echo"
)

func RunApplication() {
	// Initialize dependencies, such as database and redis connections
	properties := getProperties()
	cuslogger.SetupLogging(properties.App.Mode)

	e := echo.New()
	database := databaseConnect(properties)

	redis := ConnectCache(properties)
	defer redis.Close()

	userRepository := _userRepository.NewUserRepository(database, redis)
	userMapper := _userMapper.NewUserMapper()
	userUsecase := _userUsecase.NewUserUsecase(userRepository, userMapper)
	_userHandler.NewUserHandler(e, userUsecase)

	orderItemsRepository := _orderItemsRepository.NewOrderItemsRepository(database, redis)
	orderItemsMapper := _orderItemsMapper.NewOrderItemsMapper()
	orderItemsUsecase := _orderItemsUsecase.NewOrderItemsUsecase(orderItemsRepository, orderItemsMapper)
	_orderItemsHandler.NewOrderItemsHandler(e, orderItemsUsecase)

	orderHistoriesRepository := _orderHistoriesRepository.NewOrderHistoriesRepository(database, redis)
	orderHistoriesMapper := _orderHistoriesMapper.NewOrderHistoriesMapper()
	orderHistoriesUsecase := _orderHistoriesUsecase.NewOrderHistoriesUsecase(orderHistoriesRepository, orderHistoriesMapper)
	_orderHistoriesHandler.NewOrderHistoriesHandler(e, orderHistoriesUsecase)

	// Start the server
	e.Start(":8080")
}
