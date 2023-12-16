package db

import (
	"errors"
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var Conn *sqlx.DB

func Connect() error {
	var err error

	if os.Getenv("DATABASE_URL") == "" {
		return errors.New("DATABASE_URL not set")
	}

	Conn, err = sqlx.Connect("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return err
	}
	return nil
}

func MigrationsUp() error {
	if Conn == nil {
		return errors.New("db not initialized")
	}

	driver, err := postgres.WithInstance(Conn.DB, &postgres.Config{})

	if err != nil {
		return fmt.Errorf("error creating postgres driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file:///migrations",
		"postgres", driver)

	if err != nil {
		return fmt.Errorf("error creating migration instance: %v", err)
	}

	err = m.Up()

	if err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("error running migrations: %v", err)
	}

	return nil
}
