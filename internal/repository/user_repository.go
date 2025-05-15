package repository

import (
	"errors"
	"go-crud-msql/internal/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	FindAll() ([]models.User, error)
	FindById(id uint) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) CreateUser(user *models.User) error {
	// Check if email is changing and already used by another user
	if user.Email != "" {
		var emailCheck models.User
		if err := r.db.Where("email = ?", user.Email).First(&emailCheck).Error; err == nil {
			return errors.New("email already in use by another user")
		}
	}

	return r.db.Create(user).Error
}

func (r *userRepository) FindAll() ([]models.User, error) {
	var allUsers []models.User

	err := r.db.Find(&allUsers).Error

	return allUsers, err
}

func (r *userRepository) FindById(id uint) (*models.User, error) {
	var user models.User

	err := r.db.First(&user, id).Error

	return &user, err
}

func (r *userRepository) UpdateUser(user *models.User) error {
	// 1. Check if user exists
	var existing models.User
	if err := r.db.First(&existing, user.Id).Error; err != nil {
		return errors.New("user not found")
	}

	// 2. Check if email is changing and already used by another user
	if user.Email != existing.Email {
		var emailCheck models.User
		if err := r.db.Where("email = ?", user.Email).First(&emailCheck).Error; err == nil && emailCheck.Id != user.Id {
			return errors.New("email already in use by another user")
		}
	}

	// 3. Update fields (avoid overriding CreatedAt)
	existing.Name = user.Name
	existing.Email = user.Email
	existing.Password = user.Password
	existing.PhoneNo = user.PhoneNo

	// 4. Save updated user (GORM will auto-update UpdatedAt)
	return r.db.Save(&existing).Error
}

func (r *userRepository) DeleteUser(id uint) error {
	return r.db.Delete(&models.User{}, id).Error
}
