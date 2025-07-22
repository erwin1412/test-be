package service

import (
	"rentalapi/model"
	"rentalapi/repository"
)

type OrderService interface {
	GetAllOrders() ([]model.Order, error)
	GetOrderByID(id int) (model.Order, error)
	CreateOrder(order model.Order) (model.Order, error)
	UpdateOrder(order model.Order) (model.Order, error)
	DeleteOrder(id int) error
}
type orderService struct {
	repo repository.OrderRepository
}

func NewOrderService(repo repository.OrderRepository) OrderService {
	return &orderService{repo: repo}
}
func (s *orderService) GetAllOrders() ([]model.Order, error) {
	return s.repo.GetAll()
}
func (s *orderService) GetOrderByID(id int) (model.Order, error) {
	return s.repo.GetByID(id)
}

func (s *orderService) CreateOrder(order model.Order) (model.Order, error) {
	return s.repo.Create(order)
}
func (s *orderService) UpdateOrder(order model.Order) (model.Order, error) {
	return s.repo.Update(order)
}
func (s *orderService) DeleteOrder(id int) error {
	return s.repo.Delete(id)
}
