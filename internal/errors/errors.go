package serviceErr

type ErrorConfig struct {
	DuplicateAccountErr ErrorDetail
	InternalServerErr   ErrorDetail
	DeviceNotFound		ErrorDetail
	InvalidDeviceId		ErrorDetail
	Unauthorized		ErrorDetail
	InvalidParams		ErrorDetail
	UserAccountNotFound		ErrorDetail
	InvalidPassword		ErrorDetail
}

type ErrorDetail struct {
	Status	int
	Code    int
	Message string
}

func NewErrorConfig() *ErrorConfig {
	return &ErrorConfig{
		InvalidParams: ErrorDetail{
			Status: 400,
			Code: 400001,
			Message: "invalid params",
		},
		InvalidDeviceId: ErrorDetail{
			Status: 400,
			Code:    400002,
			Message: "invalid device id",
		},
		Unauthorized: ErrorDetail{
			Status: 401,
			Code:    401001,
			Message: "unauthorized",
		},
		InvalidPassword: ErrorDetail{
			Status: 401,
			Code:    401002,
			Message: "invalid password",
		},
		DeviceNotFound: ErrorDetail{
			Status: 404,
			Code:    404001,
			Message: "device not found",
		},
		UserAccountNotFound: ErrorDetail{
			Status: 404,
			Code:    404002,
			Message: "account not found",
		},
		DuplicateAccountErr: ErrorDetail{
			Status: 409,
			Code:    409001,
			Message: "account has already exists",
		},
		InternalServerErr: ErrorDetail{
			Status: 500,
			Code:    500,
			Message: "internal server error",
		},
	}
}