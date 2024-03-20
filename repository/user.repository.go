package repository

import (
	"golang-authen/models"

	"gorm.io/gorm"
)

type IUserRepository interface {
	Signup(user *models.User) error
	SignIn(email, password string) (*models.User, error)
}

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) Signup(user *models.User) error {

	result := r.DB.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *UserRepository) SignIn(email, password string) (*models.User, error) {
	user := &models.User{}
	result := r.DB.Where("email = ? AND password = ?", email, password).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
