package routes

import (
	"go-receiver/internal/controller"
	"go-receiver/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetDeviceRoutes(router *gin.Engine, deviceController controller.DeviceController) {
	deviceRouter := router.Group("/device")
	// deviceRouter.Use(cors.Default())
	deviceRouter.Use(middleware.AuthMiddleware())

	deviceRouter.POST("/create", deviceController.CreateSingleDevice)
	deviceRouter.GET("/:id", deviceController.GetSingleDevice)
	
}