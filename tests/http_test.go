package tests

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/fabiante/persurl/api"
	"github.com/fabiante/persurl/tests/driver"
	"github.com/fabiante/persurl/tests/specs"
	"github.com/gin-gonic/gin"
	_ "modernc.org/sqlite"
)

func TestWithHTTPDriver(t *testing.T) {
	gin.SetMode(gin.TestMode)
	handler := gin.Default()

	sqlitePath := "./mydb.sqlite"

	// ensure file does not exist
	_ = os.Remove(sqlitePath)

	db, err := sql.Open("sqlite", sqlitePath)
	if err != nil {
		panic(err)
	}

	err = migrateDb(db)
	if err != nil {
		panic(err)
	}

	server := api.NewServer(db)
	api.SetupRouting(handler, server)

	testServer := httptest.NewServer(handler)

	dr := driver.NewHTTPDriver(testServer.URL, http.DefaultTransport)

	specs.TestResolver(t, dr)
	specs.TestAdministration(t, dr)
}

func migrateDb(db *sql.DB) error {
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
