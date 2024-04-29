package controller

import (
	"go-receiver/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type usersController struct {
	usersService service.UsersService
}

type UsersController interface {
	GetSingleUser(ctx *gin.Context)
}

func NewUsersController(usersService service.UsersService) UsersController {
	return &usersController{usersService: usersService}
}

type GetSingleUserRes struct {
	Name string `json:"name"`
	Account string `json:"account"`
}
func (c *usersController) GetSingleUser(ctx *gin.Context) {
	userId := ctx.Param("id")
	
	id, parseErr := strconv.ParseUint(userId, 10, 64)
	if parseErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
	}

	user, getUserErr := c.usersService.GetUserById(uint(id))
	if getUserErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Interval error"})
		return
	}
	if user == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	userRes := &GetSingleUserRes{
		Name: user.Name,
		Account: user.Account,
	}

	ctx.JSON(http.StatusOK, userRes)
}