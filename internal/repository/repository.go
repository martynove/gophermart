package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/martynove/gophermart/internal/models"
	"github.com/sirupsen/logrus"
)

type Repository struct {
	Authorization
	logger *logrus.Logger
}

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUserByLogin(user models.User) (bool, error)
}

func NewRepository(logger *logrus.Logger, db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(logger, db),
		logger:        logger,
	}
}
