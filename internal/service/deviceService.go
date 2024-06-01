package service

import (
	"fmt"
	"go-receiver/internal/models"
	"go-receiver/internal/repository"

	"github.com/google/uuid"
)

type deviceService struct {
	deviceRepo repository.DeviceRepository
}

type DeviceService interface {
	GetDeviceById(id uint) (*models.Device, error)
	CreateOneDevice(userId uint, deviceName string) error
}

func NewDeviceService(deviceRepo repository.DeviceRepository) DeviceService {
	return &deviceService{deviceRepo: deviceRepo}
}

func (s *deviceService) GetDeviceById(id uint) (*models.Device, error) {
	device, err := s.deviceRepo.FindOneByID(id)

	return device, err
}

func (s *deviceService) CreateOneDevice(userId uint, deviceName string) error {
	newUUID := uuid.New()

	device := models.Device{
		Name: deviceName,
		UUID: newUUID.String(),
		UserId: userId,
	}

	createErr := s.deviceRepo.CreateSingleDevice(device)
	if createErr != nil {
		fmt.Println("createErr =", createErr)
		return createErr
	}

	return nil
}