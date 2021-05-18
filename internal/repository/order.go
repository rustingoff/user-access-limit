package repository

import (
	"user_access_limit/internal/models"

	"github.com/jmoiron/sqlx"
)

type OrderRepository struct {
	db *sqlx.DB
}

func NewOrderRepository(db *sqlx.DB) *OrderRepository {
	return &OrderRepository{db: db}
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
