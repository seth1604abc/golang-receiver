package middleware

import (
	"fmt"
	"go-receiver/configs"
	"go-receiver/internal/service"
	"net/http"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorizaed",
			})
			c.Abort()
			return
		}

		isValid, claims, validErr := validToken(token)
		if validErr != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "internal server error",
			})
			fmt.Println(validErr)
			c.Abort()
			return
		}

		if !isValid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorizaed",
			})
			c.Abort()
			return
		}

		userId, idErr := strconv.Atoi(claims.StandardClaims.Subject)
		if idErr != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorizaed",
			})
			c.Abort()
			return
		}

		c.Set("user", uint(userId))
		c.Next()
	}
}

func validToken(t string) (isValid bool, claims *service.Claims, err error) {
	tokenSlice := strings.Split(t, " ")
	if len(tokenSlice) != 2 {
		return false, nil, nil
	}

	tokenStr := tokenSlice[1]

	cl := &service.Claims{}

	token, err := jwt.ParseWithClaims(tokenStr, cl, func(token *jwt.Token) (interface{}, error){
		return []byte(configs.Configs.App.JWTSecret), nil
	})
	if err != nil {
		// 若是過期則不判斷為解析錯誤
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return false, nil, nil
			}
		}
		return false, nil, err
	}

	if !token.Valid{
		return false, nil, nil
	}

	return true, cl, nil
}