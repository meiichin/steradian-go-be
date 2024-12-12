package repositories

import (
	"gorm.io/gorm"
	"steradian-go/models"
)

type OrderRepository interface {
	GetAll() ([]models.Order, error)
	GetByID(id uint) (models.Order, error)
	Create(order models.Order) error
	Update(order models.Order) error
	Delete(id uint) error
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db}
}

func (o orderRepository) GetAll() ([]models.Order, error) {
	var orders []models.Order
	//err := o.db.Find(&orders).Error
	err := o.db.Preload("Car").Find(&orders).Error
	return orders, err
}

func (o orderRepository) GetByID(id uint) (models.Order, error) {
	var order models.Order
	err := o.db.Preload("Car").First(&order, id).Error
	return order, err
}

func (o orderRepository) Create(order models.Order) error {
	return o.db.Create(&order).Error
}

func (o orderRepository) Update(order models.Order) error {
	return o.db.Save(&order).Error
}

func (o orderRepository) Delete(id uint) error {
	return o.db.Delete(&models.Order{}, id).Error
}
