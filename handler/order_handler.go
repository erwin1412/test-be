package handler

import (
	"net/http"
	"rentalapi/model"
	"rentalapi/service"
	"strconv"

	"github.com/labstack/echo/v4"
)

type OrderHandler struct {
	service service.OrderService
}

func NewOrderHandler(service service.OrderService) *OrderHandler {
	return &OrderHandler{service}
}
func (h *OrderHandler) GetAll(c echo.Context) error {
	orders, err := h.service.GetAllOrders()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, orders)
}

func (h *OrderHandler) GetByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	order, err := h.service.GetOrderByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, order)
}

func (h *OrderHandler) Create(c echo.Context) error {
	var order model.Order
	if err := c.Bind(&order); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	createdOrder, err := h.service.CreateOrder(order)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, createdOrder)
}

func (h *OrderHandler) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var order model.Order
	if err := c.Bind(&order); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	order.ID = id
	updatedOrder, err := h.service.UpdateOrder(order)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, updatedOrder)
}
func (h *OrderHandler) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.service.DeleteOrder(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
