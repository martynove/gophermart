package repository

import (
	"fmt"
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
	var isExist bool
	q := fmt.Sprintf("SELECT EXISTS (SELECT 1 FROM %s WHERE login=$1)", usersTable)
	row := r.db.QueryRow(q, user.Login)
	if err := row.Scan(&isExist); err != nil {
		return id, err
	}
	if isExist {
		return 0, models.ErrorLoginExist // make const error
	}
	query := fmt.Sprintf("INSERT INTO %s (login, password_hash) values ($1, $2) RETURNING id", usersTable)
	row = r.db.QueryRow(query, user.Login, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUserByLogin(user models.User) (bool, error) {
	var hasExist bool
	return hasExist, nil
}
