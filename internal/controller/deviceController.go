package controller

import (
	"fmt"
	serviceErr "go-receiver/internal/errors"
	"go-receiver/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type deviceController struct {
	deviceService 	service.DeviceService
	serviceErr		*serviceErr.ErrorConfig
}

type DeviceController interface {
	GetSingleDevice(ctx *gin.Context)
	CreateSingleDevice(ctx *gin.Context)
}

func NewDeviceController(deviceService service.DeviceService, serviceErr *serviceErr.ErrorConfig) DeviceController {
	return &deviceController{deviceService: deviceService, serviceErr: serviceErr}
}

func (c *deviceController) GetSingleDevice(ctx *gin.Context) {
	deviceId := ctx.Param("id")
	userId, exist := ctx.Get("user")
	if !exist {
		ctx.JSON(c.serviceErr.Unauthorized.Status, gin.H{
			"data": map[string]interface{}{
					"code": c.serviceErr.Unauthorized.Code,
					"message": c.serviceErr.Unauthorized.Message,
				},
		})
		return
	}

	id, parseErr := strconv.ParseUint(deviceId, 10, 64)
	if parseErr != nil {
		fmt.Printf("parseErr is %v", parseErr)
		ctx.JSON(c.serviceErr.InvalidDeviceId.Code, gin.H{"error": c.serviceErr.InvalidDeviceId.Message})
		return
	}

	device, getUserErr := c.deviceService.GetDeviceById(uint(id))
	if getUserErr != nil {
		fmt.Printf("parseErr is %v", getUserErr)
		ctx.JSON(c.serviceErr.InternalServerErr.Code, gin.H{"error": c.serviceErr.InternalServerErr.Message})
		return
	}
	if device == nil {
		fmt.Println("device is nil")
		ctx.JSON(c.serviceErr.DeviceNotFound.Code, gin.H{"error": c.serviceErr.DeviceNotFound.Message})
		return
	}

	if device.UserId != userId {
		fmt.Printf("userId is %T", device.UserId)
		ctx.JSON(c.serviceErr.DeviceNotFound.Code, gin.H{"error": c.serviceErr.DeviceNotFound.Message})
		return
	}

	ctx.JSON(http.StatusOK, device)
}

type CreateSingleDeviceParams struct {
	DeviceName string `json:"deviceName" validate:"required"`
}
func (c *deviceController) CreateSingleDevice(ctx *gin.Context) {
	userId, exist := ctx.Get("user")
	if !exist {
		ctx.JSON(c.serviceErr.Unauthorized.Status, gin.H{
			"data": map[string]interface{}{
				"code": c.serviceErr.Unauthorized.Code,
				"message":c.serviceErr.Unauthorized.Message,
			},
		})
		return
	}

	uid, ok := userId.(uint)
	if !ok {
		ctx.JSON(c.serviceErr.Unauthorized.Status, gin.H{
			"data": map[string]interface{}{
				"code": c.serviceErr.Unauthorized.Code,
				"message":c.serviceErr.Unauthorized.Message,
			},
		})
		return
	}

	var body CreateSingleDeviceParams

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(c.serviceErr.InvalidParams.Status, gin.H{
			"data": map[string]interface{}{
				"code": c.serviceErr.InvalidParams.Code,
				"message":c.serviceErr.InvalidParams.Message,
			},
		})
		return
	}

	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		ctx.JSON(c.serviceErr.InvalidParams.Status, gin.H{
			"data": map[string]interface{}{
				"code": c.serviceErr.InvalidParams.Code,
				"message":c.serviceErr.InvalidParams.Message,
			},
		})
		return
	}

	err := c.deviceService.CreateOneDevice(uid, body.DeviceName)
	if err != nil {
		ctx.JSON(c.serviceErr.InternalServerErr.Status, gin.H{
			"data": map[string]interface{}{
				"code": c.serviceErr.InternalServerErr.Code,
				"message":c.serviceErr.InternalServerErr.Message,
			},
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": map[string]interface{}{
			"code": http.StatusOK,
			"message":"create device success",
		},
	})
}