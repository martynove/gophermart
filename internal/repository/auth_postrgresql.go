package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/martynove/gophermart/internal/models"
	"github.com/sirupsen/logrus"
)

type AuthPostgres struct {
	db     *sqlx.DB
	logger *logrus.Logger
}

func NewAuthPostgres(logger *logrus.Logger, db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db,
		logger: logger}
}

func (r *AuthPostgres) CreateUser(user models.User) (int, error) {
	var id int
	return id, nil
}

func (r *AuthPostgres) GetUserByLogin(user models.User) (bool, error) {
	var hasExist bool
	return hasExist, nil
}
