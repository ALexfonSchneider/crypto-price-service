package main

import (
	"crypto-price-service/internal/config"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

var postgresDialect = string(goose.DialectPostgres)

func main() {
	cfg := config.MustConfig()

	db, err := sql.Open("postgres", cfg.Postgres.DSN())

	if err != nil {
		panic(err)
	}

	if err = goose.SetDialect(postgresDialect); err != nil {
		panic(err)
	}

	if err := goose.Up(db, "migrations"); err != nil {
		panic(err)
	}
}
