package model

import "time"

type Order struct {
	ID              int       `json:"id"`
	CarID           int       `json:"car_id"`
	OrderDate       time.Time `json:"order_date"`
	PickupDate      time.Time `json:"pickup_date"`
	DropoffDate     time.Time `json:"dropoff_date"`
	PickupLocation  string    `json:"pickup_location"`
	DropoffLocation string    `json:"dropoff_location"`
}
