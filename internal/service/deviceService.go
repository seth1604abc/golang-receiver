package service

import (
	"go-receiver/internal/models"
	"go-receiver/internal/repository"
)

type deviceService struct {
	deviceRepo repository.DeviceRepository
}

type DeviceService interface {
	GetDeviceById(id uint) (*models.Device, error)
}

func NewDeviceService(deviceRepo repository.DeviceRepository) DeviceService {
	return &deviceService{deviceRepo: deviceRepo}
}

func (s *deviceService) GetDeviceById(id uint) (*models.Device, error) {
	device, err := s.deviceRepo.FindOneByID(id)

	return device, err
}