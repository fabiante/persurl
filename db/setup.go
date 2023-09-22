package db

import (
	"database/sql"
	"fmt"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	"github.com/fabiante/persurl/db/migrations"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Handles is a collection of database handles.
//
// The contained handles are all using the same underlying SQL connection.
type Handles struct {
	Std  *sql.DB
	Goqu *goqu.Database
	Gorm *gorm.DB
}

func SetupAndMigratePostgresDB(dsn string, maxConnections int) (Handles, error) {
	handles, err := SetupPostgresDB(dsn, maxConnections)
	if err != nil {
		return handles, err
	}

	err = migrations.RunPostgres(handles.Std)
	if err != nil {
		return handles, fmt.Errorf("migrating postgres database failed: %s", err)
	}

	return handles, nil
}

func SetupPostgresDB(dsn string, maxConnections int) (Handles, error) {
	database, err := sql.Open("postgres", dsn)
	if err != nil {
		return Handles{}, fmt.Errorf("opening postgres database failed: %s", err)
	}

	database.SetMaxOpenConns(maxConnections)

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: database,
	}))
	if err != nil {
		return Handles{}, fmt.Errorf("setting up gorm database failed: %w", err)
	}

	return Handles{
		Std:  database,
		Goqu: goqu.New("postgres", database),
		Gorm: gormDB,
	}, nil
}
