package repository

import (
	"user_access_limit/internal/models"

	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) Create(user models.User) (int, error) {
	return 0, nil
}

func (u *UserRepository) Update(userId int, user models.User) (int, error) {
	return 0, nil
}

func (u *UserRepository) Delete(userId int) error {
	return nil
}

func (u *UserRepository) View() ([]models.User, error) {
	return []models.User{}, nil
}
