package handler

import (
	"net/http"
	"rentalapi/model"
	"rentalapi/service"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CarHandler struct {
	service service.CarService
}

func NewCarHandler(service service.CarService) *CarHandler {
	return &CarHandler{service}
}

// GetAll godoc
// @Summary Get all cars
// @Description Get a list of all cars
// @Tags cars
// @Accept json
// @Produce json
// @Success 200 {array} model.Car
// @Failure 500 {string} string "Internal Server Error"
// @Router /cars [get]
func (h *CarHandler) GetAll(c echo.Context) error {
	cars, err := h.service.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, cars)
}

// GetByID godoc
// @Summary Get car by ID
// @Description Get a car by its ID
// @Tags cars
// @Accept json
// @Produce json
// @Param id path int true "Car ID"
// @Success 200 {object} model.Car
// @Failure 404 {string} string "Car not found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /cars/{id} [get]
func (h *CarHandler) GetByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	car, err := h.service.GetByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, car)
}

// Create godoc
// @Summary Create a new car
// @Description Create a new car with the provided details
// @Tags cars
// @Accept json
// @Produce json
// @Param car body model.Car true "Car details"
// @Success 201 {object} model.Car
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /cars [post]
func (h *CarHandler) Create(c echo.Context) error {
	var car model.Car
	if err := c.Bind(&car); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	createdCar, err := h.service.Create(car)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, createdCar)
}

// Update godoc
// @Summary Update an existing car
// @Description Update a car with the provided details
// @Tags cars
// @Accept json
// @Produce json
// @Param id path int true "Car ID"
// @Param car body model.Car true "Car details"
// @Success 200 {object} model.Car
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "Car not found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /cars/{id} [put]
func (h *CarHandler) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var car model.Car
	if err := c.Bind(&car); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	car.ID = id
	updatedCar, err := h.service.Update(car)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, updatedCar)
}

// Delete godoc
// @Summary Delete a car
// @Description Delete a car by its ID
// @Tags cars
// @Accept json
// @Produce json
// @Param id path int true "Car ID"
// @Success 204 "No Content"
// @Failure 404 {string} string "Car not found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /cars/{id} [delete]
func (h *CarHandler) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.service.Delete(id); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
