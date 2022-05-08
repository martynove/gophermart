package repository

import "github.com/jmoiron/sqlx"

const (
	usersTable = "users"
)

func NewPostgresDB(source string) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", source)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, err
}
