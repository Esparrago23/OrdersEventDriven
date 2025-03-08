package main

import (
	"log"
	orders "order_Event_Driven/src/orders/infraestructure"
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error al cargar el archivo .env")
	}

	allowOrigins := os.Getenv("ALLOW_ORIGINS")
	originsList := strings.Split(allowOrigins, ",")
	log.Println("ALLOW_ORIGINS:", allowOrigins)
	for _, origin := range originsList {
		if !strings.HasPrefix(origin, "http://") && !strings.HasPrefix(origin, "https://") {
			log.Fatalf("Invalid origin: %s. Origins must start with http:// or https://", origin)
		}
	}
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     originsList,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	orders.Init(r)
	r.Run(":8080")
}
