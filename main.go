package main

import (
	"time"
	"github.com/gin-contrib/cors"
	orders "order_Event_Driven/src/orders/infraestructure"
	"github.com/gin-gonic/gin"
)

func main() {
	r:= gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	orders.Init(r)
	r.Run(":8080")
}