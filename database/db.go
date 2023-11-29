package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	_ "github.com/mattes/migrate/source/file"

	"github.com/pkg/errors"
)

func New() (*sqlx.DB, error) {
	var (
		host         = os.Getenv("DB_HOST")
		port         = os.Getenv("DB_PORT")
		user         = os.Getenv("DB_USER")
		password     = os.Getenv("DB_PASSWORD")
		dbname       = os.Getenv("DB_NAME")
		migrationDir = os.Getenv("MIGRATION_DIR")
	)

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create connection to database")
	}

	if err = db.Ping(); err != nil {
		return nil, errors.Wrap(err, "database doesn't respond")
	}

	if err = runMigrations(db.DB, dbname, migrationDir); err != nil {
		return nil, errors.Wrap(err, "run migrations error")
	}

	return db, nil
}

func runMigrations(db *sql.DB, dbname, migrationDir string) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(migrationDir, dbname, driver)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}
