package repository

import (
	"rentalapi/config"
	"rentalapi/model"
)

type CarRepository interface {
	GetAll() ([]model.Car, error)
	GetByID(id int) (model.Car, error)
	Create(car model.Car) (model.Car, error)
	Update(car model.Car) (model.Car, error)
	Delete(id int) error
}
type carRepository struct{}

func NewCarRepository() CarRepository {
	return &carRepository{}
}

func (r *carRepository) GetAll() ([]model.Car, error) {
	rows, err := config.DB.Query("SELECT id, car_name, day_rate, month_rate , image FROM cars")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var cars []model.Car
	for rows.Next() {
		var car model.Car
		if err := rows.Scan(&car.ID, &car.CarName, &car.DayRate, &car.MonthRate, &car.Image); err != nil {
			return nil, err
		}
		cars = append(cars, car)
	}
	return cars, nil
}

func (r *carRepository) GetByID(id int) (model.Car, error) {
	var car model.Car
	err := config.DB.QueryRow("SELECT id, car_name, day_rate, month_rate, image FROM cars WHERE id = $1", id).Scan(&car.ID, &car.CarName, &car.DayRate, &car.MonthRate, &car.Image)
	if err != nil {
		return model.Car{}, err
	}
	return car, nil
}
func (r *carRepository) Create(car model.Car) (model.Car, error) {
	err := config.DB.QueryRow("INSERT INTO cars (car_name, day_rate, month_rate, image) VALUES ($1, $2, $3, $4) RETURNING id", car.CarName, car.DayRate, car.MonthRate, car.Image).Scan(&car.ID)
	if err != nil {
		return model.Car{}, err
	}
	return car, nil
}
func (r *carRepository) Update(car model.Car) (model.Car, error) {
	_, err := config.DB.Exec("UPDATE cars SET car_name = $1, day_rate = $2, month_rate = $3, image = $4 WHERE id = $5", car.CarName, car.DayRate, car.MonthRate, car.Image, car.ID)
	if err != nil {
		return model.Car{}, err
	}
	return car, nil
}
func (r *carRepository) Delete(id int) error {
	_, err := config.DB.Exec("DELETE FROM cars WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
