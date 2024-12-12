package models

type Order struct {
	OrderID         uint   `gorm:"primaryKey" json:"order_id"`
	CarID           uint   `gorm:"not null" json:"car_id"`
	OrderDate       string `gorm:"not null" json:"order_date"`
	DropoffDate     string `gorm:"not null" json:"dropoff_date"`
	PickupLocation  string `gorm:"not null" json:"pickup_location"`
	DropoffLocation string `gorm:"not null" json:"dropoff_location"`
	Car             Car    `gorm:"foreignKey:CarID;references:CarID"`
}
