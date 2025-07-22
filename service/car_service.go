package service

import (
	"rentalapi/model"
	"rentalapi/repository"
)

type CarService interface {
	GetAll() ([]model.Car, error)
	GetByID(id int) (model.Car, error)
	Create(car model.Car) (model.Car, error)
	Update(car model.Car) (model.Car, error)
	Delete(id int) error
}

type carService struct {
	repo repository.CarRepository
}

func NewCarService(repo repository.CarRepository) CarService {
	return &carService{repo: repo}
}

func (s *carService) GetAll() ([]model.Car, error) {
	return s.repo.GetAll()
}

func (s *carService) GetByID(id int) (model.Car, error) {
	return s.repo.GetByID(id)
}
func (s *carService) Create(car model.Car) (model.Car, error) {
	return s.repo.Create(car)
}
func (s *carService) Update(car model.Car) (model.Car, error) {
	return s.repo.Update(car)
}
func (s *carService) Delete(id int) error {
	return s.repo.Delete(id)
}
