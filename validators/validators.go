package validators

type CarsValidator struct {
	CarName   string  `json:"car_name" binding:"required"`
	DayRate   float64 `json:"day_rate" binding:"required"`
	MonthRate float64 `json:"month_rate" binding:"required"`
	Image     string  `json:"image" binding:"required,url"`
}
