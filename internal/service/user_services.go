package service

import (
	"errors"
	"go-crud-msql/internal/models"
	"go-crud-msql/internal/repository"
)

type UserService interface {
	CreateUser(user *models.User) error
	FindAll() ([]models.User, error)
	FindById(id uint) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id uint) error
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo}
}

// CreateUser implements UserService.
func (s *userService) CreateUser(user *models.User) error {
	if user.Name == "" || user.Email == "" || user.Password == "" {
		return errors.New("name, email and password are required")
	}

	return s.userRepo.CreateUser(user)
}

// FindAll implements UserService.
func (s *userService) FindAll() ([]models.User, error) {
	return s.userRepo.FindAll()
}

// FindById implements UserService.
func (s *userService) FindById(id uint) (*models.User, error) {
	return s.userRepo.FindById(id)
}

// UpdateUser implements UserService.
func (s *userService) UpdateUser(user *models.User) error {
	// Check if name, email and password are provided
	if user.Name == "" || user.Email == "" || user.Password == "" {
		return errors.New("name, email and password are required")
	}

	return s.userRepo.UpdateUser(user)
}

// DeleteUser implements UserService.
func (s *userService) DeleteUser(id uint) error {
	_, err := s.userRepo.FindById(id)
	if err != nil {
		return err
	}
	return s.userRepo.DeleteUser(id)
}
