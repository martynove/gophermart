package service

import (
	"github.com/martynove/gophermart/internal/models"
	"github.com/martynove/gophermart/internal/repository"
)

type OrderService struct {
	repo repository.Order
}

func NewOrderService(repo repository.Order) *OrderService {
	return &OrderService{repo: repo}
}

func (o *OrderService) UploadOrder(userID, orderNumber int) error {
	return o.repo.UploadOrder(userID, orderNumber)
}

func (o *OrderService) GetAllOrders(userID int) ([]models.Order, error) {
	return o.repo.GetAllOrders(userID)
}
