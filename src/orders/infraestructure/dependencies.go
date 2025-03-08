package infraestructure

import (
	"order_Event_Driven/src/orders/application"
	"order_Event_Driven/src/orders/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {
	os := NewMySQL()

	createOrderService := application.NewCreateOrderUseCase(os)
	deleteOrderService := application.NewDeleteOrderUseCase(os)
	findOrderByIdService := application.NewFindOrderByIdUseCase(os)
	findAllOrdersService := application.NewFindAllOrdersUseCase(os)
	updateOrderService := application.NewUpdateOrderUseCase(os)

	createOrderController := controllers.NewCreateOrderController(*createOrderService)
	deleteOrderController := controllers.NewDeleteOrderController(*deleteOrderService)
	findOrderByIdController := controllers.NewFindOrderByIdController(*findOrderByIdService)
	findAllOrdersController := controllers.NewFindAllOrdersController(*findAllOrdersService)
	updateOrderController := controllers.NewUpdateOrderController(*updateOrderService)

	OrdersRoutes(router, OrdersHandlers{
		Create:   createOrderController,
		Delete:   deleteOrderController,
		FindById: findOrderByIdController,
		FindAll:  findAllOrdersController,
		Update:   updateOrderController,
	})
}
