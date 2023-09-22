package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fabiante/persurl/api"
	"github.com/fabiante/persurl/app"
	"github.com/fabiante/persurl/config"
	"github.com/fabiante/persurl/db"
	"github.com/fabiante/persurl/tests/driver"
	"github.com/fabiante/persurl/tests/specs"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestWithHTTPDriver(t *testing.T) {
	gin.SetMode(gin.TestMode)
	handler := gin.Default()

	_, database, err := db.SetupAndMigratePostgresDB(config.DbDSN(), config.DbMaxConnections())
	require.NoError(t, err, "setting up db failed")

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: database,
	}))
	require.NoError(t, err, "setting up gorm db failed")

	err = db.EmptyTables(database, "purls", "domains")
	require.NoError(t, err, "truncating tables failed")

	service := db.NewDatabase(database)
	server := api.NewServer(service, app.NewAdminService(gormDB))
	api.SetupRouting(handler, server)

	testServer := httptest.NewServer(handler)

	dr := driver.NewHTTPDriver(testServer.URL, http.DefaultTransport)

	specs.TestResolver(t, dr)
	specs.TestAdministration(t, dr)
}
