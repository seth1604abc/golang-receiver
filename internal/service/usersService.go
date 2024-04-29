package service

import (
	"go-receiver/internal/models"
	"go-receiver/internal/repository"
)

type usersService struct {
	usersRepo repository.UsersRepository
}

type UsersService interface {
	GetUserById(id uint) (*models.Users, error)
}

func NewUsersService(userRepo repository.UsersRepository) UsersService {
	return &usersService{usersRepo: userRepo}
}

func (s *usersService) GetUserById(id uint) (*models.Users, error) {
	user, err := s.usersRepo.FindOneByID(id)

	return user, err
}