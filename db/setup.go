package db

import (
	"database/sql"
	"fmt"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/sqlite3"
	"github.com/fabiante/persurl/db/migrations"
	_ "modernc.org/sqlite"
)

func SetupAndMigrateDB(path string) (*sql.DB, *goqu.Database, error) {
	db, gdb, err := SetupDB(path)
	if err != nil {
		return nil, nil, err
	}

	err = migrations.Run(db)
	if err != nil {
		return nil, nil, fmt.Errorf("migrating sqlite database failed: %s", err)
	}

	return db, gdb, nil
}

func SetupDB(path string) (*sql.DB, *goqu.Database, error) {
	database, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, nil, fmt.Errorf("opening sqlite database failed: %s", err)
	}

	return database, goqu.New("sqlite3", database), nil
}
