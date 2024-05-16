package routes

import (
	"go-receiver/internal/controller"

	"github.com/gin-gonic/gin"
)

func SetAuthRoutes(router *gin.Engine, authController controller.AuthController) {
	authRouter := router.Group("/auth")

	authRouter.POST("/register", authController.RegisterUser)
	authRouter.POST("/login", authController.LoginUser)
}