package repository

import (
	"user_access_limit/internal/models"

	"github.com/jmoiron/sqlx"
)

type User interface {
	Create(user models.User) (int, error)
	Update(userId int, user models.User) (int, error)
	Delete(userId int) error
	View() ([]models.User, error)
}

type Order interface {
	Create(order models.Order) (int, error)
	Update(orderId int, order models.Order) (int, error)
	Delete(orderId int) error
	View() ([]models.Order, error)
}

type Customer interface {
	Create(custom models.Customer) (int, error)
	Update(customerId int, customer models.Customer) (int, error)
	Delete(customerId int) error
	View() ([]models.Customer, error)
}

type Repository struct {
	User
	Order
	Customer
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User:     NewUserRepository(db),
		Order:    NewOrderRepository(db),
		Customer: NewCustomerRepository(db),
	}
}
