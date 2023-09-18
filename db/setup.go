package db

import (
	"database/sql"
	"fmt"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	"github.com/fabiante/persurl/db/migrations"
	_ "github.com/lib/pq"
)

func SetupAndMigratePostgresDB(dsn string) (*sql.DB, *goqu.Database, error) {
	db, gdb, err := SetupPostgresDB(dsn)
	if err != nil {
		return nil, nil, err
	}

	err = migrations.RunPostgres(db)
	if err != nil {
		return nil, nil, fmt.Errorf("migrating postgres database failed: %s", err)
	}

	return db, gdb, nil
}

func SetupPostgresDB(dsn string) (*sql.DB, *goqu.Database, error) {
	database, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, nil, fmt.Errorf("opening postgres database failed: %s", err)
	}

	return database, goqu.New("postgres", database), nil
}
