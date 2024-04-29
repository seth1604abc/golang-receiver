package routes

import (
	"go-receiver/internal/controller"

	"github.com/gin-gonic/gin"
)

func SetDeviceRoutes(router *gin.Engine, deviceController controller.DeviceController) {
	deviceRouter := router.Group("/device")

	deviceRouter.GET("/:id", deviceController.GetSingleDevice)
}