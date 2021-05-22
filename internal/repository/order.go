package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/rustingoff/user-access-limit/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepository struct {
	mongo *mongo.Client
	pg    *gorm.DB
}

func NewOrderRepository(mongo *mongo.Client, pg *gorm.DB) *OrderRepository {
	return &OrderRepository{mongo, pg}
}

func (u *OrderRepository) Create(order models.Order) (int, error) {
	return 0, nil
}

func (u *OrderRepository) Update(orderId int, order models.Order) (int, error) {
	return 0, nil
}

func (u *OrderRepository) Delete(orderId int) error {
	return nil
}

func (u *OrderRepository) View() ([]models.Order, error) {
	return []models.Order{}, nil
}
