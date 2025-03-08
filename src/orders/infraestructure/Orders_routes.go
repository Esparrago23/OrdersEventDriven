package infraestructure

import (
	"order_Event_Driven/src/orders/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

type OrdersHandlers struct {
	Create   *controllers.CreateOrderController
	Delete   *controllers.DeleteOrderController
	FindById *controllers.FindOrderByIdController
	FindAll  *controllers.FindAllOrdersController
	Update   *controllers.UpdateOrderController
	Payment  *controllers.ProccessPaymentController
}

func OrdersRoutes(router *gin.Engine, handlers OrdersHandlers) {
	ordersGroup := router.Group("/orders")
	{
		ordersGroup.POST("/", handlers.Create.Execute)
		ordersGroup.DELETE("/:id", handlers.Delete.Execute)
		ordersGroup.GET("/:id", handlers.FindById.Execute)
		ordersGroup.GET("/", handlers.FindAll.Execute)
		ordersGroup.PUT("/:id", handlers.Update.Execute)
		ordersGroup.POST("/:id/pay", handlers.Payment.Execute)
	}
}
