package main

import (
	"rentalapi/config"
	"rentalapi/handler"
	"rentalapi/repository"
	"rentalapi/service"

	_ "rentalapi/docs"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Rental Mobil API
// @version 1.0
// @description This is a simple API for managing car rentals.
// @host localhost:8082
// @BasePath /
func main() {
	// Initialize Echo framework
	e := echo.New()

	config.ConnectDB()

	// Initialize services and handlers
	carRepo := repository.NewCarRepository()
	carService := service.NewCarService(carRepo)
	carHandler := handler.NewCarHandler(carService)
	orderRepo := repository.NewOrderRepository()
	orderService := service.NewOrderService(orderRepo)
	orderHandler := handler.NewOrderHandler(orderService)

	// Define routes for car operations
	e.GET("/cars", carHandler.GetAll)
	e.GET("/cars/:id", carHandler.GetByID)
	e.POST("/cars", carHandler.Create)
	e.PUT("/cars/:id", carHandler.Update)
	e.DELETE("/cars/:id", carHandler.Delete)
	// Define routes for order operations
	e.GET("/orders", orderHandler.GetAll)
	e.GET("/orders/:id", orderHandler.GetByID)
	e.POST("/orders", orderHandler.Create)
	e.PUT("/orders/:id", orderHandler.Update)
	e.DELETE("/orders/:id", orderHandler.Delete)
	// Swagger documentation route
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	// Start the server
	e.Logger.Fatal(e.Start(":8082"))
}
