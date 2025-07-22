package service_test

import (
	"rentalapi/model"
	"rentalapi/repository"
	"rentalapi/service"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock Repo
type MockOrderRepo struct {
	mock.Mock
	repository.OrderRepository
}

func (m *MockOrderRepo) IsCarAvailable(carID int, pickupDate, dropoffDate time.Time) (bool, error) {
	args := m.Called(carID, pickupDate, dropoffDate)
	return args.Bool(0), args.Error(1)
}

func (m *MockOrderRepo) Create(order model.Order) (model.Order, error) {
	args := m.Called(order)
	return args.Get(0).(model.Order), args.Error(1)
}

func (m *MockOrderRepo) GetAll() ([]model.Order, error) {
	args := m.Called()
	return args.Get(0).([]model.Order), args.Error(1)
}

func (m *MockOrderRepo) GetByID(id int) (model.Order, error) {
	args := m.Called(id)
	return args.Get(0).(model.Order), args.Error(1)
}

func (m *MockOrderRepo) Update(order model.Order) (model.Order, error) {
	args := m.Called(order)
	return args.Get(0).(model.Order), args.Error(1)
}

func (m *MockOrderRepo) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestCreateOrder_Valid(t *testing.T) {
	mockRepo := new(MockOrderRepo)
	svc := service.NewOrderService(mockRepo)

	order := model.Order{
		CarID:           1,
		OrderDate:       time.Now(),
		PickupDate:      time.Now().AddDate(0, 0, 1),
		DropoffDate:     time.Now().AddDate(0, 0, 2),
		PickupLocation:  "Jakarta",
		DropoffLocation: "Bandung",
	}

	mockRepo.On("IsCarAvailable", order.CarID, order.PickupDate, order.DropoffDate).Return(true, nil)
	mockRepo.On("Create", order).Return(order, nil)

	result, err := svc.CreateOrder(order)

	assert.NoError(t, err)
	assert.Equal(t, order, result)
	mockRepo.AssertExpectations(t)
}

func TestCreateOrder_InvalidDates(t *testing.T) {
	mockRepo := new(MockOrderRepo)
	svc := service.NewOrderService(mockRepo)

	now := time.Now()

	// pickup_date < order_date
	order := model.Order{
		CarID:       1,
		OrderDate:   now,
		PickupDate:  now.AddDate(0, 0, -1),
		DropoffDate: now.AddDate(0, 0, 2),
	}
	_, err := svc.CreateOrder(order)
	assert.EqualError(t, err, "pickup_date must be same or after order_date")

	// dropoff_date < pickup_date
	order = model.Order{
		CarID:       1,
		OrderDate:   now,
		PickupDate:  now.AddDate(0, 0, 2),
		DropoffDate: now.AddDate(0, 0, 1),
	}
	_, err = svc.CreateOrder(order)
	assert.EqualError(t, err, "dropoff_date must be same or after pickup_date")
}

func TestCreateOrder_Overlap(t *testing.T) {
	mockRepo := new(MockOrderRepo)
	svc := service.NewOrderService(mockRepo)

	now := time.Now()

	order := model.Order{
		CarID:       1,
		OrderDate:   now,
		PickupDate:  now.AddDate(0, 0, 1),
		DropoffDate: now.AddDate(0, 0, 3),
	}

	mockRepo.On("IsCarAvailable", order.CarID, order.PickupDate, order.DropoffDate).Return(false, nil)

	_, err := svc.CreateOrder(order)
	assert.EqualError(t, err, "car is not available for the selected dates")
	mockRepo.AssertExpectations(t)
}

func TestGetAllOrders(t *testing.T) {
	mockRepo := new(MockOrderRepo)
	svc := service.NewOrderService(mockRepo)

	expected := []model.Order{
		{ID: 1}, {ID: 2},
	}

	mockRepo.On("GetAll").Return(expected, nil)

	result, err := svc.GetAllOrders()

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

func TestGetOrderByID(t *testing.T) {
	mockRepo := new(MockOrderRepo)
	svc := service.NewOrderService(mockRepo)

	expected := model.Order{ID: 123}

	mockRepo.On("GetByID", 123).Return(expected, nil)

	result, err := svc.GetOrderByID(123)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

func TestUpdateOrder(t *testing.T) {
	mockRepo := new(MockOrderRepo)
	svc := service.NewOrderService(mockRepo)

	order := model.Order{ID: 10}

	mockRepo.On("Update", order).Return(order, nil)

	result, err := svc.UpdateOrder(order)

	assert.NoError(t, err)
	assert.Equal(t, order, result)
	mockRepo.AssertExpectations(t)
}

func TestDeleteOrder(t *testing.T) {
	mockRepo := new(MockOrderRepo)
	svc := service.NewOrderService(mockRepo)

	mockRepo.On("Delete", 1).Return(nil)

	err := svc.DeleteOrder(1)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
