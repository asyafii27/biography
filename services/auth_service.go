package services

import (
	"biography-api/helpers"
	"biography-api/models"
	"biography-api/repositories"
	"errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(name, email, password string) (*models.User, error)
	Login(email, password string) (string, error)
}

type authService struct {
	userRepo repositories.UserRepository
}

func NewAuthService(userRepo repositories.UserRepository) AuthService {
	return &authService{userRepo}
}

func (s *authService) Register(name, email, password string) (*models.User, error) {
	// Check if user already exists
	_, err := s.userRepo.GetUserByEmail(email)
	if err == nil {
		return nil, errors.New("email already registered")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		ID:       uuid.New(),
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
	}

	err = s.userRepo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *authService) Login(email, password string) (string, error) {
	user, err := s.userRepo.GetUserByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := helpers.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
