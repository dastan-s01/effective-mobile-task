package db

import (
	"context"
	"errors"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"os"
)

func PgConnection() (*pgxpool.Pool, error) {
	dsn := os.Getenv("DATABASE_URL")

	if dsn == "" {
		log.Fatal("DATABASE_URL не установлен в .env или пуст.")
	}

	dbPool, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		log.Println("Не удается подключиться к БД", err)
		return nil, err
	}

	m, err := migrate.New("file://db/migrations", dsn)
	if err != nil {
		return nil, err
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return nil, err
	}

	return dbPool, nil
}
