package service

import (
	"github.com/trucktrace/internal/models"
	"github.com/trucktrace/internal/repository"
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

type Services struct {
	User
	Order
	Customer
}

func NewService(repos *repository.Repository) *Services {
	return &Services{
		User:     NewUserService(repos.User),
		Order:    NewOrderService(repos.Order),
		Customer: NewCustomerService(repos.Customer),
	}
}
