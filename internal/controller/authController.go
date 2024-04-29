package controller

import (
	"go-receiver/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type authController struct {
	authService service.AuthService
}

type AuthController interface {
	RegisterUser(ctx *gin.Context)
}

func NewAuthController(authService service.AuthService) AuthController {
	return &authController{authService: authService}
}

type RegisterUserReq struct {
	Account  string `validate:"required" json:"account"`
	Password string `validate:"required" json:"password"`
	Name     string `validate:"required" json:"name"`
	Email    string `validate:"required,email" json:"email"`
}
func (s *authController) RegisterUser(ctx *gin.Context) {
	// 檢查參數
	var body RegisterUserReq
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// 驗證參數
	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	params := service.RegisterUserParams{
		Account: body.Account,
		Password: body.Password,
		Name: body.Name,
		Email: body.Email,
	}
	userErr := s.authService.RegisterUser(params)
	if userErr != nil {
		ctx.JSON(userErr.Code, gin.H{"message": userErr.Message})
		return
	}

	ctx.JSON(200, gin.H{"message": "register success"})
}