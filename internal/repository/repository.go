package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/trucktrace/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
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

func NewRepository(mongo *mongo.Client, pg *gorm.DB) *Repository {
	return &Repository{
		User:     NewUserRepository(mongo, pg),
		Order:    NewOrderRepository(mongo, pg),
		Customer: NewCustomerRepository(mongo, pg),
	}
}
