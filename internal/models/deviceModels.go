package models

import (
	"time"

	"gorm.io/gorm"
)

type Device struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey;column:id"`
	ConnectString  string `gorm:"column:connect_string;not null"`
	UserId  uint `gorm:"column:user_id;not null"`
	CreatedAt *time.Time `gorm:"column:created_at;type:timestamp"`
	UpdatedAt *time.Time `gorm:"column:updated_at;type:timestamp"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp"`
	
	User Users `gorm:"foreignKey:UserId"`
}

func (Device) TableName() string {
	return "device"
}