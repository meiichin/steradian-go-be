package repositories

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"

	"steradian-go/models"
)

func setupMockDB() (*gorm.DB, sqlmock.Sqlmock, error) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	// Expect GORM's initialization queries
	mock.ExpectQuery("SELECT VERSION()").WillReturnRows(sqlmock.NewRows([]string{"VERSION()"}).AddRow("5.7.34"))

	dialector := mysql.New(mysql.Config{
		DriverName: "mysql",
		DSN:        "sqlmock_db_0",
		Conn:       mockDB,
	})

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}

	return db, mock, nil
}
func TestCarRepository_GetAll(t *testing.T) {
	db, mock, err := setupMockDB()
	assert.NoError(t, err)
	repo := NewCarRepository(db)

	// Define expected results
	expectedCars := []models.Car{
		{CarID: 1, CarName: "Car 1", DayRate: 0, MonthRate: 0, Image: ""},
		{CarID: 2, CarName: "Car 2", DayRate: 0, MonthRate: 0, Image: ""},
	}

	// Mock database rows
	rows := sqlmock.NewRows([]string{"car_id", "car_name", "day_rate", "month_rate", "image"}).
		AddRow(expectedCars[0].CarID, expectedCars[0].CarName, expectedCars[0].DayRate, expectedCars[0].MonthRate, expectedCars[0].Image).
		AddRow(expectedCars[1].CarID, expectedCars[1].CarName, expectedCars[1].DayRate, expectedCars[1].MonthRate, expectedCars[1].Image)

	// Expect query with backticks
	mock.ExpectQuery("SELECT \\* FROM `cars`").WillReturnRows(rows)

	// Call the function
	cars, err := repo.GetAll()

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, expectedCars, cars)
	assert.NoError(t, mock.ExpectationsWereMet())
}
