// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"go-receiver/internal/controller"
	"go-receiver/internal/database"
	"go-receiver/internal/errors"
	"go-receiver/internal/repository"
	"go-receiver/internal/service"
)

// Injectors from wire.go:

func InitialUsersController() (controller.UsersController, error) {
	db, err := database.GetDB()
	if err != nil {
		return nil, err
	}
	usersRepository, err := repository.NewUsersRepository(db)
	if err != nil {
		return nil, err
	}
	usersService := service.NewUsersService(usersRepository)
	usersController := controller.NewUsersController(usersService)
	return usersController, nil
}

func InitializeAuthController() (controller.AuthController, error) {
	db, err := database.GetDB()
	if err != nil {
		return nil, err
	}
	usersRepository, err := repository.NewUsersRepository(db)
	if err != nil {
		return nil, err
	}
	errorConfig := serviceErr.NewErrorConfig()
	authService := service.NewAuthService(usersRepository, errorConfig)
	authController := controller.NewAuthController(authService)
	return authController, nil
}

func InitializeDeviceController() (controller.DeviceController, error) {
	db, err := database.GetDB()
	if err != nil {
		return nil, err
	}
	deviceRepository, err := repository.NewDeviceRepository(db)
	if err != nil {
		return nil, err
	}
	deviceService := service.NewDeviceService(deviceRepository)
	errorConfig := serviceErr.NewErrorConfig()
	deviceController := controller.NewDeviceController(deviceService, errorConfig)
	return deviceController, nil
}
