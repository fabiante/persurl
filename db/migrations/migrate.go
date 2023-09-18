package migrations

import (
	"database/sql"
	"fmt"
)

func Run(db *sql.DB) error {
	stmts := migrationsSQLite

	for i, stmt := range stmts {
		_, err := db.Exec(stmt)
		if err != nil {
			return fmt.Errorf("stmts[%d] failed: %w", i, err)
		}
	}

	return nil
}

var migrationsSQLite = []string{
	`create table main.domains
(
    id   integer      not null
        constraint domains_pk
            primary key autoincrement,
    name varchar(128) not null
        constraint domains_pk2
            unique
)
`,
	`
create table purls
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
);`,
}
