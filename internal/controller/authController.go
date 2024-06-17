package controller

import (
	serviceErr "go-receiver/internal/errors"
	"go-receiver/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type authController struct {
	authService 	service.AuthService
	serviceError	*serviceErr.ErrorConfig
}

type AuthController interface {
	RegisterUser(ctx *gin.Context)
	LoginUser(ctx *gin.Context)
}

func NewAuthController(authService service.AuthService, serviceError *serviceErr.ErrorConfig) AuthController {
	return &authController{authService: authService, serviceError: serviceError}
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
		ctx.JSON(s.serviceError.InvalidParams.Status, gin.H{
			"data": map[string]interface{}{
				"code": s.serviceError.InvalidParams.Code,
				"message": s.serviceError.InvalidParams.Message,
			}})
		return
	}

	// 驗證參數
	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		ctx.JSON(s.serviceError.InvalidParams.Status, gin.H{
			"data": map[string]interface{}{
				"code": s.serviceError.InvalidParams.Code,
				"message": s.serviceError.InvalidParams.Message,
			}})
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
		ctx.JSON(userErr.Status, gin.H{
			"data": map[string]interface{}{
				"code": userErr.Code,
				"message": userErr.Message,
			}})
		return
	}

	ctx.JSON(200, gin.H{
		"data": map[string]interface{}{
			"code": 200,
			"message": "register success",
		}})
}

type LoginUserParams struct {
	Account string `validate:"required" json:"account"`
	Password string `validate:"required" json:"password"`
}
func (s *authController) LoginUser(ctx *gin.Context) {
	var body LoginUserParams

	// 檢查參數
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(s.serviceError.InvalidParams.Status, gin.H{
			"data": map[string]interface{}{
				"code": s.serviceError.InvalidParams.Code,
				"message": s.serviceError.InvalidParams.Message,
			}})
		return
	}

	// 驗證參數
	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		ctx.JSON(s.serviceError.InvalidParams.Status, gin.H{
			"data": map[string]interface{}{
				"code": s.serviceError.InvalidParams.Code,
				"message": s.serviceError.InvalidParams.Message,
			}})
		return
	}

	token, tokenErr := s.authService.LoginUser(service.LoginUserParams{Account: body.Account, Password: body.Password})
	if tokenErr != nil {
		ctx.JSON(tokenErr.Status, gin.H{
			"data": map[string]interface{}{
				"code": tokenErr.Code,
				"message": tokenErr.Message,
			}})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": map[string]interface{}{
			"code": http.StatusOK,
			"message": "login success",
			"accessToken": token,
	}})
}