package services

import (
	"errors"
	"steradian-go/models"
	"steradian-go/repository"
)

//go:generate mockery --name=CarService --filename car_service_mock.go --outpkg services
type CarService interface {
	GetAllCars() ([]models.Car, error)
	GetCarByID(id uint) (models.Car, error)
	CreateCar(car models.Car) error
	UpdateCar(id uint, car models.Car) error
	DeleteCar(id uint) error
}

type carService struct {
	repo repositories.CarRepository
}

func NewCarService(repo repositories.CarRepository) CarService {
	return &carService{repo}
}

func (s *carService) GetAllCars() ([]models.Car, error) {
	return s.repo.GetAll()
}

func (s *carService) GetCarByID(id uint) (models.Car, error) {
	car, err := s.repo.GetByID(id)
	if err != nil {
		return car, errors.New("car not found")
	}
	return car, nil
}

func (s *carService) CreateCar(car models.Car) error {
	return s.repo.Create(car)
}

func (s *carService) UpdateCar(id uint, car models.Car) error {
	_, err := s.repo.GetByID(id)
	if err != nil {
		return errors.New("car not found")
	}
	car.CarID = id
	return s.repo.Update(car)
}

func (s *carService) DeleteCar(id uint) error {
	_, err := s.repo.GetByID(id)
	if err != nil {
		return errors.New("car not found")
	}
	return s.repo.Delete(id)
}
