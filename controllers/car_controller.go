package controllers

import (
	"net/http"
	"steradian-go/utils"
	"steradian-go/validators"
	"strconv"

	"github.com/gin-gonic/gin"
	"steradian-go/models"
	"steradian-go/services"
)

type CarController struct {
	service services.CarService
}

func NewCarController(service services.CarService) *CarController {
	return &CarController{service}
}

func (c *CarController) GetAllCars(ctx *gin.Context) {
	cars, err := c.service.GetAllCars()
	if err != nil {
		response := utils.APIResponse(false, err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}
	response := utils.APIResponse(true, "Get Cars Success", cars)
	ctx.JSON(http.StatusOK, response)
}

func (c *CarController) GetCarByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	car, err := c.service.GetCarByID(uint(id))
	if err != nil {
		errors := utils.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := utils.APIResponse(false, "Car Not Found", errorMessage)
		ctx.JSON(http.StatusNotFound, response)
		return
	}
	ctx.JSON(http.StatusOK, car)
}

func (c *CarController) CreateCar(ctx *gin.Context) {
	var newCar validators.CarsValidator
	if err := ctx.ShouldBindJSON(&newCar); err != nil {
		errors := utils.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := utils.APIResponse(false, "Add new Car Failed", errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	car := models.Car{
		CarName:   newCar.CarName,
		DayRate:   newCar.DayRate,
		MonthRate: newCar.MonthRate,
		Image:     newCar.Image,
	}

	if err := c.service.CreateCar(car); err != nil {
		errors := utils.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := utils.APIResponse(false, "Add Car Failed", errorMessage)
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}
	response := utils.APIResponse(true, "Add new Car Success", car)
	ctx.JSON(http.StatusCreated, response)
}

func (c *CarController) UpdateCar(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var newCar validators.CarsValidator
	if err := ctx.ShouldBindJSON(&newCar); err != nil {
		errors := utils.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := utils.APIResponse(false, "Add new Car Failed", errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	// check car found
	_, err := c.service.GetCarByID(uint(id))
	if err != nil {
		response := utils.APIResponse(false, "Invalid car ID", nil)
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	car := models.Car{
		CarName:   newCar.CarName,
		DayRate:   newCar.DayRate,
		MonthRate: newCar.MonthRate,
		Image:     newCar.Image,
	}
	if err := c.service.UpdateCar(uint(id), car); err != nil {
		errors := utils.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := utils.APIResponse(false, "Update Car Failed", errorMessage)
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	response := utils.APIResponse(true, "Update Car Success", car)
	ctx.JSON(http.StatusCreated, response)
}

func (c *CarController) DeleteCar(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	_, err := c.service.GetCarByID(uint(id))
	if err != nil {
		response := utils.APIResponse(false, "Invalid car ID", nil)
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	if err := c.service.DeleteCar(uint(id)); err != nil {
		errors := utils.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := utils.APIResponse(false, "Update Car Failed", errorMessage)
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	response := utils.APIResponse(true, "Delete Cake Success", nil)
	ctx.JSON(http.StatusOK, response)
}
