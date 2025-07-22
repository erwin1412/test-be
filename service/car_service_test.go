package service_test

import (
	"rentalapi/model"
)

type mockCarRepository struct {
	mockCreate func(car model.Car) error
}

func (m *mockCarRepository) GetAll() ([]model.Car, error) {
	return []model.Car{
		{ID: 1, CarName: "Car A", DayRate: 100000, MonthRate: 2500000, Image: "imageA.jpg"},
		{ID: 2, CarName: "Car B", DayRate: 120000, MonthRate: 2800000, Image: "imageB.jpg"},
	}, nil
}
func (m *mockCarRepository) GetByID(id int) (model.Car, error) {
	if id == 1 {
		return model.Car{ID: 1, CarName: "Car A", DayRate: 100000, MonthRate: 2500000, Image: "imageA.jpg"}, nil
	} else if id == 2 {
		return model.Car{ID: 2, CarName: "Car B", DayRate: 120000, MonthRate: 2800000, Image: "imageB.jpg"}, nil
	}
	return model.Car{}, nil
}
func (m *mockCarRepository) Create(car model.Car) (model.Car, error) {
	car.ID = 3
	return car, nil
}
func (m *mockCarRepository) Update(car model.Car) (model.Car, error) {
	if car.ID == 1 {
		car.ID = 1
		return car, nil
	}
	return model.Car{}, nil
}
func (m *mockCarRepository) Delete(id int) error {
	if id == 1 {
		return nil
	}
	return nil
}
