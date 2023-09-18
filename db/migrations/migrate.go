package migrations

import (
	"database/sql"
	"fmt"

	"github.com/lopezator/migrator"
)

func RunSQLite(db *sql.DB) error {
	// Configure migrations

	m, err := migrator.New(
		// type cast is required because []*migrator.MigrationNoTx is not assignable to []migrator.Migration
		migrator.Migrations(migrationsSQLite...),
	)
	if err != nil {
		return fmt.Errorf("initializing migrations failed: %w", err)
	}

	// Migrate up
	if err := m.Migrate(db); err != nil {
		return fmt.Errorf("running migrations failed: %w", err)
	}

	return nil
}

func RunPostgres(db *sql.DB) error {
	// Configure migrations

	m, err := migrator.New(
		// type cast is required because []*migrator.MigrationNoTx is not assignable to []migrator.Migration
		migrator.Migrations(migrationsPostgres...),
	)
	if err != nil {
		return fmt.Errorf("initializing migrations failed: %w", err)
	}

	// Migrate up
	if err := m.Migrate(db); err != nil {
		return fmt.Errorf("running migrations failed: %w", err)
	}

	return nil
}

func newMigration(name string, query string) *migrator.MigrationNoTx {
	return &migrator.MigrationNoTx{
		Name: name,
		Func: func(db *sql.DB) error {
			if _, err := db.Exec(query); err != nil {
				return err
			}
			return nil
		},
	}
}

var migrationsSQLite = []any{
	newMigration("2023-09-18-00000001-CreateTableDomains", `create table main.domains
(
    id   integer      not null
        constraint domains_pk
            primary key autoincrement,
    name varchar(128) not null
        constraint domains_pk2
            unique
)`,
	),
	newMigration("2023-09-18-00000002-CreateTablePurls", `create table purls
(
    id        integer       not null
        constraint puls_pk
            primary key autoincrement,
    domain_id integer       not null
        constraint purls_domains_id_fk
            references domains
            on delete restrict,
    name      varchar(128)  not null,
    target    varchar(4096) not null,
    constraint purls_pk
        unique (domain_id, name)
)`,
	),
}

var migrationsPostgres = []any{
	newMigration("2023-09-18-00000001-CreateTableDomains", `create table domains
(
    id   serial
        constraint domains_pk2
            unique,
    name varchar(128) not null
        constraint domains_pk
            primary key
)`,
	),
	newMigration("2023-09-18-00000002-CreateTablePurls", `create table purls
(
    id        serial
        constraint purls_pk
            primary key,
    domain_id integer       not null
        constraint purls_domains_id_fk
            references domains (id)
			on delete restrict,
    name      varchar(128)  not null,
    target    varchar(4096) not null,
    constraint purls_pk2
        unique (domain_id, name)
)`,
	),
}
