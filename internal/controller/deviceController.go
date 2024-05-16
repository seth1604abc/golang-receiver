package controller

import (
	"fmt"
	serviceErr "go-receiver/internal/errors"
	"go-receiver/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type deviceController struct {
	deviceService 	service.DeviceService
	serviceErr		serviceErr.ErrorConfig
}

type DeviceController interface {
	GetSingleDevice(ctx *gin.Context)
}

func NewDeviceController(deviceService service.DeviceService) DeviceController {
	return &deviceController{deviceService: deviceService}
}

func (c *deviceController) GetSingleDevice(ctx *gin.Context) {
	deviceId := ctx.Param("id")
	userId, exist := ctx.Get("user")
	if !exist {
		ctx.JSON(c.serviceErr.Unauthorized.Code, gin.H{"error": c.serviceErr.Unauthorized.Message})
		return
	}
	fmt.Println(userId)

	id, parseErr := strconv.ParseUint(deviceId, 10, 64)
	if parseErr != nil {
		ctx.JSON(c.serviceErr.InvalidDeviceId.Code, gin.H{"error": c.serviceErr.InvalidDeviceId.Message})
		return
	}

	device, getUserErr := c.deviceService.GetDeviceById(uint(id))
	if getUserErr != nil {
		ctx.JSON(c.serviceErr.InternalServerErr.Code, gin.H{"error": c.serviceErr.InternalServerErr.Message})
		return
	}
	if device == nil {
		ctx.JSON(c.serviceErr.DeviceNotFound.Code, gin.H{"error": c.serviceErr.DeviceNotFound.Message})
		return
	}

	if device.UserId != userId {
		ctx.JSON(c.serviceErr.DeviceNotFound.Code, gin.H{"error": c.serviceErr.DeviceNotFound.Message})
		return
	}

	ctx.JSON(http.StatusOK, device)
}