package main

import (
	"fmt"
	"log"

	// MQTT "github.com/eclipse/paho.mqtt.golang"
	"go-receiver/configs"

	"github.com/gin-gonic/gin"

	"go-receiver/internal/database"
	"go-receiver/internal/routes"

	"github.com/gin-contrib/cors"
)

func main() {
	_, dbErr := database.GetDB()
	if dbErr != nil {
		log.Fatalf("Database init error: %v", dbErr)
	}

	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8080"}

	router.Use(cors.New(config))

	routes.SetUpRouter(router)

	router.Run(":3000")

	fmt.Printf("Server is running at %s:%d", configs.Configs.App.Host, configs.Configs.App.Port)
}