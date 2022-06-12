package repository

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	"log"
)

const (
	usersTable = "users"
)

func NewPostgresDB(source string) (*sqlx.DB, error) {
	log.Println(source)
	if err := migrateDB(source); err != nil {
		return nil, err
	}
	db, err := sqlx.Open("postgres", source)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, err
}

func migrateDB(source string) error {
	m, err := migrate.New("file://migrations/schema", source)
	if err != nil {
		return fmt.Errorf("cannot init DB migrations: %w", err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("cannot apply migrations: %w", err)
	}
	return nil
}
