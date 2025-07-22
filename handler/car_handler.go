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
func (h *CarHandler) GetAll(c echo.Context) error {
	cars, err := h.service.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, cars)
}

func (h *CarHandler) GetByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	car, err := h.service.GetByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, car)
}

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
func (h *CarHandler) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.service.Delete(id); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
