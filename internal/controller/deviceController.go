package controller

import (
	"go-receiver/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type deviceController struct {
	deviceService service.DeviceService
}

type DeviceController interface {
	GetSingleDevice(ctx *gin.Context)
}

func NewDeviceController(deviceService service.DeviceService) DeviceController {
	return &deviceController{deviceService: deviceService}
}

func (c *deviceController) GetSingleDevice(ctx *gin.Context) {
	userId := ctx.Param("id")
	
	id, parseErr := strconv.ParseUint(userId, 10, 64)
	if parseErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid device ID"})
	}

	device, getUserErr := c.deviceService.GetDeviceById(uint(id))
	if getUserErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Interval error"})
		return
	}
	if device == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Device not found"})
		return
	}

	ctx.JSON(http.StatusOK, device)
}