package service

import (
	"golang-authen/models"
	"golang-authen/repository"
)

type UserService struct {
	repo repository.IUserRepository
}

type IUserService interface {
	SignIn(email, password string) (*models.User, error)
	Signup(user *models.User) error
}

func (s *UserService) SignIn(email, password string) (*models.User, error) {
	return s.repo.SignIn(email, password)
}

func (s *UserService) Signup(user *models.User) error {
	return s.repo.Signup(user)
}
