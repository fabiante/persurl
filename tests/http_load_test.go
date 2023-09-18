package tests

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/fabiante/persurl/api"
	"github.com/fabiante/persurl/db"
	"github.com/fabiante/persurl/tests/driver"
	"github.com/fabiante/persurl/tests/specs"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestLoadWithHTTPDriver(t *testing.T) {
	if os.Getenv("TEST_LOAD") == "" {
		t.Skip("load tests are skipped because TEST_LOAD env variable is not set")
	}

	gin.SetMode(gin.TestMode)
	handler := gin.Default()

	sqlitePath := "./test_load_http.sqlite"
	_ = os.Remove(sqlitePath) // remove to ensure a clean database
	_, database, err := db.SetupAndMigrateDB(sqlitePath)
	require.NoError(t, err, "setting up db failed")

	service := db.NewDatabase(database)
	server := api.NewServer(service)
	api.SetupRouting(handler, server)

	testServer := httptest.NewServer(handler)

	dr := driver.NewHTTPDriver(testServer.URL, http.DefaultTransport)

	specs.TestLoad(t, dr)
}
