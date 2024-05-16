package serviceErr

import "fmt"

type ErrorConfig struct {
	DuplicateAccountErr ErrorDetail
	InternalServerErr   ErrorDetail
	DeviceNotFound		ErrorDetail
	InvalidDeviceId		ErrorDetail
	Unauthorized		ErrorDetail
}

type ErrorDetail struct {
	Code    int
	Message string
}

func NewErrorConfig() *ErrorConfig {
	fmt.Println("err config")
	return &ErrorConfig{
		InvalidDeviceId: ErrorDetail{
			Code:    400,
			Message: "invalid device id",
		},
		Unauthorized: ErrorDetail{
			Code:    401,
			Message: "unauthorized",
		},
		DeviceNotFound: ErrorDetail{
			Code:    404,
			Message: "device not found",
		},
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