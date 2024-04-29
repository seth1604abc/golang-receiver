package repository

import (
	"errors"
	"fmt"
	"go-receiver/internal/database"
	"go-receiver/internal/models"

	"gorm.io/gorm"
)

type usersRepo struct {
	db *gorm.DB
}

type UsersRepository interface {
	FindOneByID(id uint) (*models.Users, error)
	FindOneByAccount(account string) (*models.Users, error)
	InsertOne(p InsertOneUserParams) (*models.Users, error)
}

func NewUsersRepository(*gorm.DB) (UsersRepository, error) {
	db, dbErr := database.GetDB()
	fmt.Println("user repo init")
	if dbErr != nil {
		return nil, dbErr
	}

	return &usersRepo{db: db}, nil
}

func (r *usersRepo) FindOneByID(id uint) (*models.Users, error) {
	user := &models.Users{}

	if err := r.db.First(user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}

func (r *usersRepo) FindOneByAccount(account string) (*models.Users, error) {
	user := &models.Users{}

	if err := r.db.Where("account = ?", account).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}

type InsertOneUserParams struct {
	Account  string
	Password string
	Name     string
	Email    string
	UUID	 string
}
func (r *usersRepo) InsertOne(p InsertOneUserParams) (*models.Users, error) {
	user := models.Users{Account: p.Account, Password: p.Password, Name: p.Name, Email: p.Email, UUID: p.UUID}

	createUser := r.db.Create(&user)
	if createUser.Error != nil {
		return nil, errors.New("user insertOne error")
	}

	return &user, nil
}