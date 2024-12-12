package models

type Car struct {
	CarID     uint    `gorm:"primaryKey" json:"car_id"`
	CarName   string  `gorm:"size:50;not null" json:"car_name"`
	DayRate   float64 `gorm:"not null" json:"day_rate"`
	MonthRate float64 `gorm:"not null" json:"month_rate"`
	Image     string  `gorm:"size:255;not null" json:"image"`
}
