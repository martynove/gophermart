package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/martynove/gophermart/internal/models"
	"github.com/sirupsen/logrus"
)

type Repository struct {
	Authorization
	Order
	logger *logrus.Logger
}

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(login, password string) (models.User, error)
}
type Order interface {
	UploadOrder(userID, orderNumber int) error
	GetAllOrders(userID int) ([]models.Order, error)
}

func NewRepository(logger *logrus.Logger, db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(logger, db),
		Order:         NewOrderPostgres(logger, db),
		logger:        logger,
	}
}
