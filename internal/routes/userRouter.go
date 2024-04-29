package routes

import (
	"go-receiver/internal/controller"

	"github.com/gin-gonic/gin"
)

func SetUsersRoutes(router *gin.Engine, userController controller.UsersController) {
	userRouter := router.Group("/users")

	userRouter.GET("/:id", userController.GetSingleUser)
}