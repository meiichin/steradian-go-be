package controllers

import (
	"net/http"
	"steradian-go/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"steradian-go/models"
	"steradian-go/services"
)

type OrderController struct {
	service services.OrderService
}

func NewOrderController(service services.OrderService) *OrderController {
	return &OrderController{service}
}

func (c *OrderController) GetAllOrders(ctx *gin.Context) {
	orders, err := c.service.GetAllOrders()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	response := utils.APIResponse(true, "Get Order Success", orders)
	ctx.JSON(http.StatusOK, response)
}

func (c *OrderController) GetOrderByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	order, err := c.service.GetOrderByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	response := utils.APIResponse(true, "Detail Order Success", order)
	ctx.JSON(http.StatusOK, response)
}

func (c *OrderController) CreateOrder(ctx *gin.Context) {
	var order models.Order
	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.service.CreateOrder(order); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	response := utils.APIResponse(true, "Create Order Success", "")
	ctx.JSON(http.StatusCreated, response)
}

func (c *OrderController) UpdateOrder(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var order models.Order
	if err := ctx.ShouldBindJSON(&order); err != nil {
		response := utils.APIResponse(true, "Update Order Failed", "")
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	if err := c.service.UpdateOrder(uint(id), order); err != nil {
		response := utils.APIResponse(true, "Update Order Failed", "")
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	response := utils.APIResponse(true, "Update Order Success", "")
	ctx.JSON(http.StatusOK, response)
}

func (c *OrderController) DeleteOrder(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := c.service.DeleteOrder(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	response := utils.APIResponse(true, "Delete Cake Success", nil)
	ctx.JSON(http.StatusOK, response)
}
