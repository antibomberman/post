package services

import (
	"antibomberman/post/internal/config"
	"antibomberman/post/internal/models"
	"antibomberman/post/internal/repositories/postgres"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UserService interface {
	Register(data models.UserCreate) error
	Login(data models.LoginRequest) (string, error)
}

type userService struct {
	cfg            *config.Config
	userRepository postgres.UserRepository
}

func NewUserService(userRepo *postgres.UserRepository, cfg *config.Config) UserService {
	return &userService{
		cfg:            cfg,
		userRepository: *userRepo,
	}
}

func (u *userService) Register(data models.UserCreate) error {
	exists, err := u.userRepository.ExistsUserByEmail(data.Email)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("user already exists")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	data.Password = string(hashedPassword)

	err = u.userRepository.CreateUser(data)
	if err != nil {
		return err
	}
	return nil
}

func (u *userService) Login(data models.LoginRequest) (string, error) {
	user, err := u.userRepository.GetUserByEmail(data.Email)
	if err != nil {
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))
	if err != nil {

	}

	// Create JWT token
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenString, err := token.SignedString([]byte(u.cfg.JWTSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
