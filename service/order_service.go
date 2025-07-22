package service

import (
	"errors"
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
	isAvailable, err := s.repo.IsCarAvailable(order.CarID, order.PickupDate.Format("2006-01-02"), order.DropoffDate.Format("2006-01-02"))
	if err != nil {
		return model.Order{}, err
	}
	if !isAvailable {
		return model.Order{}, errors.New("car is not available for the selected dates")
	}
	return s.repo.Create(order)
}
func (s *orderService) UpdateOrder(order model.Order) (model.Order, error) {
	return s.repo.Update(order)
}
func (s *orderService) DeleteOrder(id int) error {
	return s.repo.Delete(id)
}
