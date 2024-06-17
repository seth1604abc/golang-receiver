package routes

import (
	di "go-receiver/internal/init"
	"log"

	"github.com/gin-gonic/gin"
)

func SetUpRouter(router *gin.Engine) {
	userController, userControllerErr := di.InitialUsersController()
	if userControllerErr != nil {
		log.Fatal("userController init failed")
	}
	SetUsersRoutes(router, userController)

	deviceController, deviceControllerErr := di.InitializeDeviceController()
	if deviceControllerErr != nil {
		log.Fatal("deviceController init failed")
	}
	SetDeviceRoutes(router, deviceController)

	authController, authControllerErr := di.InitializeAuthController()
	if authControllerErr != nil {
		log.Fatal("authController init failed")
	}
	SetAuthRoutes(router, authController)
}