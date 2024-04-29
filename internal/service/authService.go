package service

import (
	"fmt"
	serviceErr "go-receiver/internal/errors"
	"go-receiver/internal/repository"

	"github.com/google/uuid"

	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	usersRepo repository.UsersRepository
	serviceError *serviceErr.ErrorConfig
}

type AuthService interface {
	RegisterUser(p RegisterUserParams) *serviceErr.ErrorDetail
}

func NewAuthService(usersRepo repository.UsersRepository, serviceErrConfig *serviceErr.ErrorConfig) AuthService {
	return &authService{usersRepo:usersRepo, serviceError: serviceErrConfig}
}

type RegisterUserParams struct {
	Account  string
	Password string
	Name     string
	Email    string
}
func (s *authService) RegisterUser(p RegisterUserParams) *serviceErr.ErrorDetail {
	// 檢查帳號是否重複
	user, userErr := s.usersRepo.FindOneByAccount(p.Account)
	if userErr != nil {
		fmt.Println(userErr)
		return &s.serviceError.InternalServerErr
	}
	if user != nil {
		return &s.serviceError.DuplicateAccountErr
	}
	
	// hash password
	hashedPassword, hashErr := bcrypt.GenerateFromPassword([]byte(p.Password), bcrypt.DefaultCost)
	if hashErr != nil {
		fmt.Println(hashErr)
		return &s.serviceError.InternalServerErr
	}

	// create uuid
	newUUID := uuid.New()
	uuidString := newUUID.String()

	// 建立user
	insertOneParams := repository.InsertOneUserParams{
		Account: p.Account,
		Password: string(hashedPassword),
		Email: p.Email,
		Name: p.Name,
		UUID: uuidString,
	}
	createUser, createUserErr := s.usersRepo.InsertOne(insertOneParams)
	if createUserErr != nil {
		fmt.Println(createUser)
		return &s.serviceError.InternalServerErr
	}

	return nil
}