package service

import (
	"fmt"
	"go-receiver/configs"
	serviceErr "go-receiver/internal/errors"
	"go-receiver/internal/repository"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"

	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	usersRepo repository.UsersRepository
	serviceError *serviceErr.ErrorConfig
}

type AuthService interface {
	RegisterUser(p RegisterUserParams) *serviceErr.ErrorDetail
	LoginUser(params LoginUser) (token string, err *serviceErr.ErrorDetail)
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

type Claims struct {
	Account string `json:"account"`
	Role string `json:"role"`
	jwt.StandardClaims
}
type GenerateTokenParams struct {
	Account 	string
	UserName 	string
	UserId		uint
}
func (s *authService) GenerateToken(params GenerateTokenParams) (string, *serviceErr.ErrorDetail) {
	now := time.Now()
	jwtId := params.Account + strconv.FormatInt(now.Unix(), 10)

	claims := Claims{
		Account: params.Account,
		Role: "user",
		StandardClaims: jwt.StandardClaims{
			Audience: params.Account,
			ExpiresAt: now.Add(86400 * time.Second).Unix(),
			IssuedAt: now.Unix(),
			NotBefore: now.Unix(),
			Id: jwtId,
			Issuer: "MQTT-Service",
			Subject: strconv.Itoa(int(params.UserId)),
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, tokenErr := tokenClaims.SignedString([]byte(configs.Configs.App.JWTSecret))
	if tokenErr != nil {
		fmt.Println(tokenErr)
		return "", &s.serviceError.InternalServerErr
	}

	return token, nil
}

type LoginUser struct {
	Account string
	Password string
}
func (s *authService) LoginUser(params LoginUser) (token string, err *serviceErr.ErrorDetail) {
	// find user
	user, userErr := s.usersRepo.FindOneByAccount(params.Account)
	if userErr != nil {
		fmt.Println("userErr is", userErr)
		return "", &s.serviceError.InternalServerErr
	}

	// compare password
	compareErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(params.Password))
	if compareErr != nil {
		fmt.Println("compareErr is", compareErr)
		return "", &s.serviceError.InternalServerErr
	}
	
	accessToken, tokenErr := s.GenerateToken(GenerateTokenParams{Account: params.Account, UserName: user.Name, UserId: user.ID})
	if tokenErr != nil {
		fmt.Println("tokenErr is", tokenErr)
		return "", &s.serviceError.InternalServerErr
	}

	return accessToken, nil
}