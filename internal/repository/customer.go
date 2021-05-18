package repository

import (
	"user_access_limit/internal/models"

	"github.com/jmoiron/sqlx"
)

type CustomerRepository struct {
	db *sqlx.DB
}

func NewCustomerRepository(db *sqlx.DB) *CustomerRepository {
	return &CustomerRepository{db: db}
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
