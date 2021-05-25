package service

import (
	"github.com/trucktrace/internal/models"
	"github.com/trucktrace/internal/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (u *UserService) Create(user models.User) (int, error) {
	return 0, nil
}

func (u *UserService) Update(userId int, user models.User) (int, error) {
	return 0, nil
}

func (u *UserService) Delete(userId int) error {
	return nil
}

func (u *UserService) View() ([]models.User, error) {
	return []models.User{}, nil
}
