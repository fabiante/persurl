package db

import "database/sql"

func MigrateDb(db *sql.DB) error {
	stmts := []string{
		`create table purls
(
    id     integer       not null
        constraint puls_pk
            primary key autoincrement,
    domain varchar(32)   not null,
    name   varchar(128)  not null,
    target varchar(4096) not null,
    constraint puls_pk2
        unique (name, domain)
);

create index puls_domain_index
    on purls (domain);

`,
	}

	for _, stmt := range stmts {
		_, err := db.Exec(stmt)
		if err != nil {
			return err
		}
	}

	return nil
}
