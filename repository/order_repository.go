package repository

import (
	"rentalapi/config"
	"rentalapi/model"
	"time"
)

type OrderRepository interface {
	GetAll() ([]model.Order, error)
	GetByID(id int) (model.Order, error)
	Create(order model.Order) (model.Order, error)
	Update(order model.Order) (model.Order, error)
	Delete(id int) error
	IsCarAvailable(carID int, pickupDate, dropoffDate time.Time) (bool, error)
}
type orderRepository struct{}

func NewOrderRepository() OrderRepository {
	return &orderRepository{}
}
func (r *orderRepository) GetAll() ([]model.Order, error) {
	rows, err := config.DB.Query("SELECT id, car_id, order_date, pickup_date, dropoff_date, pickup_location, dropoff_location FROM orders")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var orders []model.Order
	for rows.Next() {
		var order model.Order
		if err := rows.Scan(&order.ID, &order.CarID, &order.OrderDate, &order.PickupDate, &order.DropoffDate, &order.PickupLocation, &order.DropoffLocation); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}
func (r *orderRepository) GetByID(id int) (model.Order, error) {
	var order model.Order
	err := config.DB.QueryRow("SELECT id, car_id, order_date, pickup_date, dropoff_date, pickup_location, dropoff_location FROM orders WHERE id = $1", id).Scan(&order.ID, &order.CarID, &order.OrderDate, &order.PickupDate, &order.DropoffDate, &order.PickupLocation, &order.DropoffLocation)
	if err != nil {
		return model.Order{}, err
	}
	return order, nil
}

func (r *orderRepository) IsCarAvailable(
	carID int,
	pickupDate, dropoffDate time.Time,
) (bool, error) {

	var count int
	err := config.DB.QueryRow(`
		SELECT COUNT(*)
		FROM orders
		WHERE car_id = $1 
		AND NOT ($3 < pickup_date OR $2 > dropoff_date)
	`,
		carID, pickupDate, dropoffDate).Scan(&count)

	if err != nil {
		return false, err
	}

	return count == 0, nil
}

func (r *orderRepository) Create(order model.Order) (model.Order, error) {

	err := config.DB.QueryRow("INSERT INTO orders (car_id, order_date, pickup_date, dropoff_date, pickup_location, dropoff_location) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id", order.CarID, order.OrderDate, order.PickupDate, order.DropoffDate, order.PickupLocation, order.DropoffLocation).Scan(&order.ID)
	if err != nil {
		return model.Order{}, err
	}
	return order, nil
}
func (r *orderRepository) Update(order model.Order) (model.Order, error) {
	_, err := config.DB.Exec("UPDATE orders SET car_id = $1, order_date = $2, pickup_date = $3, dropoff_date = $4, pickup_location = $5, dropoff_location = $6 WHERE id = $7", order.CarID, order.OrderDate, order.PickupDate, order.DropoffDate, order.PickupLocation, order.DropoffLocation, order.ID)
	if err != nil {
		return model.Order{}, err
	}
	return order, nil
}
func (r *orderRepository) Delete(id int) error {
	_, err := config.DB.Exec("DELETE FROM orders WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
