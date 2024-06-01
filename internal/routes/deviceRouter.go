package routes

import (
	"go-receiver/internal/controller"
	"go-receiver/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetDeviceRoutes(router *gin.Engine, deviceController controller.DeviceController) {
	deviceRouter := router.Group("/device")
	deviceRouter.Use(middleware.AuthMiddleware())

	deviceRouter.GET("/:id", deviceController.GetSingleDevice)
	deviceRouter.POST("/", deviceController.CreateSingleDevice)
}