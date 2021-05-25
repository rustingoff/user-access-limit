package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/trucktrace/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type CustomerRepository struct {
	mongo *mongo.Client
	pg    *gorm.DB
}

func NewCustomerRepository(mongo *mongo.Client, pg *gorm.DB) *CustomerRepository {
	return &CustomerRepository{mongo, pg}
}

func (u *CustomerRepository) Create(customer models.Customer) (int, error) {
	return 0, nil
}

func (u *CustomerRepository) Update(customerId int, customer models.Customer) (int, error) {
	return 0, nil
}

func (u *CustomerRepository) Delete(customerId int) error {
	return nil
}

func (u *CustomerRepository) View() ([]models.Customer, error) {
	return []models.Customer{}, nil
}
