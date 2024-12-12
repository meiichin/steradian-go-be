package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"steradian-go/config"
	"steradian-go/controllers"
	"steradian-go/models"
	repositories "steradian-go/repository"
	"steradian-go/services"
	"strings"
)

var (
	version string
)

const (
	basePath = "/steradian"
)

func main() {
	// Set version and details
	version = readVersion()
	versionData := map[string]interface{}{
		"version": version,
	}
	config.ConnectDB()
	db := config.DB
	// Run migration
	err := config.DB.AutoMigrate(&models.Car{})
	if err != nil {
		log.Fatalf("Error during migration: %v", err)
	}

	// Migrasi tabel dengan urutan yang benar
	err = config.DB.AutoMigrate(&models.Order{})
	if err != nil {
		log.Fatalf("Error during migration (Order): %v", err)
	}

	carRepo := repositories.NewCarRepository(db)
	carService := services.NewCarService(carRepo)
	carController := controllers.NewCarController(carService)

	orderRepo := repositories.NewOrderRepository(db)
	orderService := services.NewOrderService(orderRepo)
	orderController := controllers.NewOrderController(orderService)

	router := gin.Default()
	router.GET(basePath+"/version", func(c *gin.Context) {
		c.JSON(http.StatusOK, versionData)
	})

	// Grouped Routes
	api := router.Group("/steradian/api/v1")
	{
		api.GET("/cars", carController.GetAllCars)
		api.POST("/cars", carController.CreateCar)
		api.GET("car/:id", carController.GetCarByID)
		api.PUT("car/:id", carController.UpdateCar)
		api.DELETE("car/:id", carController.DeleteCar)

		api.GET("/orders", orderController.GetAllOrders)
		api.POST("/orders", orderController.CreateOrder)
		api.GET("order/:id", orderController.GetOrderByID)
		api.PUT("order/:id", orderController.UpdateOrder)
		api.DELETE("order/:id", orderController.DeleteOrder)
	}

	router.Run(":8080")
}

// readVersion reads the version string from a JSON file and returns it.
func readVersion() string {
	type version struct {
		Version string `json:"version"`
	}
	var v version
	_ = json.Unmarshal([]byte(loadString("version.json")), &v)
	return v.Version
}

// loadString reads and returns the content of a file as a string.
func loadString(path string) string {
	content, err := os.ReadFile(path)
	if err != nil {
		return err.Error()
	}
	return strings.TrimSpace(string(content))
}
