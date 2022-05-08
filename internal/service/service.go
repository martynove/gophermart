package service

import (
	"github.com/martynove/gophermart/internal/models"
	"github.com/martynove/gophermart/internal/repository"
	"github.com/sirupsen/logrus"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUserByLogin(user models.User) (bool, error)
}

type Service struct {
	Authorization
	logger *logrus.Logger
}

func NewService(r *repository.Repository, logger *logrus.Logger) *Service {
	return &Service{
		logger:        logger,
		Authorization: NewAuthService(r.Authorization),
	}
}
