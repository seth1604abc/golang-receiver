package serviceErr

import "fmt"

type ErrorConfig struct {
	DuplicateAccountErr ErrorDetail
	InternalServerErr   ErrorDetail
}

type ErrorDetail struct {
	Code    int
	Message string
}

func NewErrorConfig() *ErrorConfig {
	fmt.Println("err config")
	return &ErrorConfig{
		DuplicateAccountErr: ErrorDetail{
			Code:    409,
			Message: "account has already exists",
		},
		InternalServerErr: ErrorDetail{
			Code:    500,
			Message: "internal server error",
		},
	}
}