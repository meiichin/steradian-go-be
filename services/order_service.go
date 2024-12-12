package services

import (
	"errors"
	"steradian-go/models"
	"steradian-go/repository"
)

type OrderService interface {
	GetAllOrders() ([]models.Order, error)
	GetOrderByID(id uint) (models.Order, error)
	CreateOrder(order models.Order) error
	UpdateOrder(id uint, order models.Order) error
	DeleteOrder(id uint) error
}

type orderService struct {
	repo repositories.OrderRepository
}

func NewOrderService(repo repositories.OrderRepository) OrderService {
	return &orderService{repo}
}

func (o *orderService) GetAllOrders() ([]models.Order, error) {
	return o.repo.GetAll()
}

func (o *orderService) GetOrderByID(id uint) (models.Order, error) {
	order, err := o.repo.GetByID(id)
	if err != nil {
		return order, errors.New("order not found")
	}
	return order, nil
}

func (o *orderService) CreateOrder(order models.Order) error {
	return o.repo.Create(order)
}

func (s *orderService) UpdateOrder(id uint, order models.Order) error {
	_, err := s.repo.GetByID(id)
	if err != nil {
		return errors.New("order not found")
	}
	order.OrderID = id
	return s.repo.Update(order)
}

func (s *orderService) DeleteOrder(id uint) error {
	_, err := s.repo.GetByID(id)
	if err != nil {
		return errors.New("order not found")
	}
	return s.repo.Delete(id)
}
