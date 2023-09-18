package tests

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/doug-martin/goqu/v9"
	"github.com/fabiante/persurl/api"
	"github.com/fabiante/persurl/config"
	"github.com/fabiante/persurl/db"
	"github.com/fabiante/persurl/tests/driver"
	"github.com/fabiante/persurl/tests/specs"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestWithHTTPDriver(t *testing.T) {
	config.LoadEnv()

	gin.SetMode(gin.TestMode)
	handler := gin.Default()

	_, database, err := db.SetupAndMigratePostgresDB(config.DbDSN())
	require.NoError(t, err, "setting up db failed")

	err = emptyTables(database, "purls", "domains")
	require.NoError(t, err, "truncating tables failed")

	service := db.NewDatabase(database)
	server := api.NewServer(service)
	api.SetupRouting(handler, server)

	testServer := httptest.NewServer(handler)

	dr := driver.NewHTTPDriver(testServer.URL, http.DefaultTransport)

	specs.TestResolver(t, dr)
	specs.TestAdministration(t, dr)
}

func emptyTables(db *goqu.Database, tables ...string) error {
	var errs []error

	for _, table := range tables {
		_, err := db.Delete(table).Executor().Exec()
		errs = append(errs, err)
	}

	return errors.Join(errs...)
}
