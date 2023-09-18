package tests

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	_ "github.com/doug-martin/goqu/v9/dialect/sqlite3"
	"github.com/fabiante/persurl/api"
	"github.com/fabiante/persurl/db"
	"github.com/fabiante/persurl/tests/driver"
	"github.com/fabiante/persurl/tests/specs"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestWithHTTPDriver(t *testing.T) {
	gin.SetMode(gin.TestMode)
	handler := gin.Default()

	sqlitePath := "./test_http.sqlite"
	_ = os.Remove(sqlitePath) // remove to ensure a clean database
	_, database, err := db.SetupDB(sqlitePath)
	require.NoError(t, err, "setting up db failed")

	service := db.NewDatabase(database)
	server := api.NewServer(service)
	api.SetupRouting(handler, server)

	testServer := httptest.NewServer(handler)

	dr := driver.NewHTTPDriver(testServer.URL, http.DefaultTransport)

	specs.TestResolver(t, dr)
	specs.TestAdministration(t, dr)
}
