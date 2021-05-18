package service

import (
	"github.com/rustingoff/user-access-limit/internal/models"
	"github.com/rustingoff/user-access-limit/internal/repository"
)

type CustomerService struct {
	repo repository.Customer
}

func NewCustomerService(repo repository.Customer) *CustomerService {
	return &CustomerService{
		repo: repo,
	}
}
func (c *CustomerService) Create(customer models.Customer) (int, error) {
	return 0, nil
}

func (c *CustomerService) Update(customerId int, customer models.Customer) (int, error) {
	return 0, nil
}

func (c *CustomerService) Delete(customerId int) error {
	return nil
}

func (u *CustomerService) View() ([]models.Customer, error) {
	return []models.Customer{}, nil
}
