package database

import (
	"fmt"

	"go-receiver/configs"

	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DBConnect *gorm.DB
	err error
	once sync.Once
)

func GetDB() (*gorm.DB, error) {
	once.Do(func () {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", configs.Configs.Db.Username, configs.Configs.Db.Password, configs.Configs.Db.Host, configs.Configs.Db.Port, configs.Configs.Db.Name)

		DBConnect, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

		if err != nil {
			return
		}
	})

	return DBConnect, err
}