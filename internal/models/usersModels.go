package models

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey;column:id"`
	Account  string `gorm:"column:account;not null;index"`
	Password  string `gorm:"column:password;type:varchar(512);not null"`
	Name  string `gorm:"column:name;not null"`
	Email  string `gorm:"column:email;not null"`
	UUID  string `gorm:"column:uuid;not null"`
	CreatedAt *time.Time `gorm:"column:created_at;type:timestamp"`
	UpdatedAt *time.Time `gorm:"column:updated_at;type:timestamp"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp"`
}

func (Users) TableName() string {
	return "users"
}