package repository

import (
	"fmt"
	"go-receiver/internal/database"
	"go-receiver/internal/models"

	"gorm.io/gorm"
)

type deviceRepo struct {
	db *gorm.DB
}

type DeviceRepository interface {
	FindOneByID(id uint) (*models.Device, error)
}

func NewDeviceRepository(*gorm.DB) (DeviceRepository, error) {
	db, dbErr := database.GetDB()
	fmt.Println("user repo init")
	if dbErr != nil {
		return nil, dbErr
	}

	return &deviceRepo{db: db}, nil
}

func (r *deviceRepo) FindOneByID(id uint) (*models.Device, error) {
	device := &models.Device{}

	if err := r.db.Joins("left join users on device.user_id = users.id").First(device, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return device, nil
}