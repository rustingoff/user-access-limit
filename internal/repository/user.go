package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/rustingoff/user-access-limit/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	mongo *mongo.Client
	pg    *gorm.DB
}

func NewUserRepository(mongo *mongo.Client, pg *gorm.DB) *UserRepository {
	return &UserRepository{mongo, pg}
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
