package controllers

import (
	"order_Event_Driven/src/orders/application"
	"net/http"
	"strconv"
	
	"order_Event_Driven/src/orders/infraestructure/rabbitmq"
	"github.com/gin-gonic/gin"
)

type ProccessPaymentController struct {
	ProcessPaymentUseCase application.ProcessPaymentUseCase
}

func NewProccessPaymentOrderController(processPaymentUseCase application.ProcessPaymentUseCase) *ProccessPaymentController {
	return &ProccessPaymentController{ProcessPaymentUseCase: processPaymentUseCase}
}

func (controller *ProccessPaymentController) Execute(c *gin.Context) {
	orderIDParam := c.Param("id")
	orderID, err := strconv.Atoi(orderIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	order, err := controller.ProcessPaymentUseCase.Execute(orderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	rabbitmq.PublishOrderPaymentStatus(order.OrderID, order.TotalAmount,order.Status)

	c.JSON(http.StatusOK, gin.H{"order": order})
}
