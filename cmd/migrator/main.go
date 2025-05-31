package main

import (
	"dice_roll__v1_not_provablyfair/internal/config"
	"errors"
	"flag"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	cfg := config.MustLoad()
	var migrationsPath, migrationsTable string
	flag.StringVar(&migrationsPath, "migrations-path", "", "path to migrations")
	flag.StringVar(&migrationsTable, "migrations-table", "migrations", "name of migrations table")
	flag.Parse()
	if migrationsPath == "" {
		panic("migrations-path is required")
	}

	sourceURL := "file://" + migrationsPath
	databaseURL := fmt.Sprintf(cfg.PostgresConnStr, migrationsTable)
	m, err := migrate.New(sourceURL, databaseURL)
	if err != nil {
		panic(err)
	}
	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("no migrations to apply")
			return
		}
		panic(err)
	}
}
