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

// GetAll godoc
// @Summary Get all orders
// @Description Get a list of all orders
// @Tags orders
// @Accept json
// @Produce json
// @Success 200 {array} model.Order
// @Failure 500 {string} string "Internal Server Error"
// @Router /orders [get]
func (h *OrderHandler) GetAll(c echo.Context) error {
	orders, err := h.service.GetAllOrders()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, orders)
}

// GetByID godoc
// @Summary Get order by ID
// @Description Get an order by its ID
// @Tags orders
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Success 200 {object} model.Order
// @Failure 404 {string} string "Order not found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /orders/{id} [get]
func (h *OrderHandler) GetByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	order, err := h.service.GetOrderByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, order)
}

// Create godoc
// @Summary Create a new order
// @Description Create a new order with the provided details
// @Tags orders
// @Accept json
// @Produce json
// @Param order body model.Order true "Order details"
// @Success 201 {object} model.Order
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /orders [post]
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

// Update godoc
// @Summary Update an existing order
// @Description Update an order with the provided details
// @Tags orders
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Param order body model.Order true "Order details"
// @Success 200 {object} model.Order
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "Order not found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /orders/{id} [put]
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

// Delete godoc
// @Summary Delete an order
// @Description Delete an order by its ID
// @Tags orders
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Success 204 {string} string "No Content"
// @Failure 404 {string} string "Order not found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /orders/{id} [delete]
func (h *OrderHandler) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.service.DeleteOrder(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
