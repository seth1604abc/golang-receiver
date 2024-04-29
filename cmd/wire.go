package di

import (
	"go-receiver/internal/controller"
	"go-receiver/internal/database"
	serviceErr "go-receiver/internal/errors"
	"go-receiver/internal/repository"
	"go-receiver/internal/service"

	"github.com/google/wire"
)

func InitialUsersController() (controller.UsersController, error) {
	wire.Build(controller.NewUsersController, service.NewUsersService, repository.NewUsersRepository, database.GetDB)
	return nil, nil
}

func InitializeAuthController() (controller.AuthController, error) {
	wire.Build(controller.NewAuthController, service.NewAuthService, serviceErr.NewErrorConfig, repository.NewUsersRepository, database.GetDB)
	return nil, nil
}

func InitializeDeviceController() (controller.DeviceController, error) {
	wire.Build(controller.NewDeviceController, service.NewDeviceService, repository.NewDeviceRepository, database.GetDB)
	return nil, nil
}