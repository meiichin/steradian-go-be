package repositories

import (
	"gorm.io/gorm"
	"steradian-go/models"
)

//go:generate mockery --name=CarRepository --filename car_repository_mock.go --outpkg repositories
type CarRepository interface {
	GetAll() ([]models.Car, error)
	GetByID(id uint) (models.Car, error)
	Create(car models.Car) error
	Update(car models.Car) error
	Delete(id uint) error
}

type carRepository struct {
	db *gorm.DB
}

func NewCarRepository(db *gorm.DB) CarRepository {
	return &carRepository{db}
}

func (r *carRepository) GetAll() ([]models.Car, error) {
	var cars []models.Car
	err := r.db.Find(&cars).Error
	return cars, err
}

func (r *carRepository) GetByID(id uint) (models.Car, error) {
	var car models.Car
	err := r.db.First(&car, id).Error
	return car, err
}

func (r *carRepository) Create(car models.Car) error {
	return r.db.Create(&car).Error
}

func (r *carRepository) Update(car models.Car) error {
	return r.db.Save(&car).Error
}

func (r *carRepository) Delete(id uint) error {
	return r.db.Delete(&models.Car{}, id).Error
}
