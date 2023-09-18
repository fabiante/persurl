package db

import (
	"database/sql"
	"fmt"

	"github.com/fabiante/persurl/db/migrations"
	_ "modernc.org/sqlite"
)

func SetupDB(path string) (*sql.DB, error) {
	database, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, fmt.Errorf("opening sqlite database failed: %s", err)
	}

	err = migrations.Run(database)
	if err != nil {
		return nil, fmt.Errorf("migrating sqlite database failed: %s", err)
	}

	return database, nil
}
