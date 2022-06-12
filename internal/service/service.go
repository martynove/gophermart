package service

import (
	"github.com/martynove/gophermart/internal/models"
	"github.com/martynove/gophermart/internal/repository"
	"github.com/sirupsen/logrus"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(login, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}

type Order interface {
	UploadOrder(userID, orderNumber int) error
	GetAllOrders(userID int) ([]models.Order, error)
}

type Service struct {
	Authorization
	Order
	logger *logrus.Logger
}

func NewService(r *repository.Repository, logger *logrus.Logger) *Service {
	return &Service{
		logger:        logger,
		Authorization: NewAuthService(r.Authorization),
		Order:         NewOrderService(r.Order),
	}
}
