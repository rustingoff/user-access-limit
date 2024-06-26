package service

import (
	"github.com/trucktrace/internal/models"
	"github.com/trucktrace/internal/repository"
)

type OrderService struct {
	repo repository.Order
}

func NewOrderService(repo repository.Order) *OrderService {
	return &OrderService{
		repo: repo,
	}
}

func (o *OrderService) Create(order models.Order) (int, error) {
	return 0, nil
}

func (o *OrderService) Update(orderId int, order models.Order) (int, error) {
	return 0, nil
}

func (o *OrderService) Delete(orderId int) error {
	return nil
}

func (o *OrderService) View() ([]models.Order, error) {
	return []models.Order{}, nil
}
