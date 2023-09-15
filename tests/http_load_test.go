package tests

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/doug-martin/goqu/v9"
	"github.com/fabiante/persurl/api"
	"github.com/fabiante/persurl/db"
	"github.com/fabiante/persurl/tests/driver"
	"github.com/fabiante/persurl/tests/specs"
	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"
	"github.com/stretchr/testify/require"
)

func TestLoadWithHTTPDriver(t *testing.T) {
	if os.Getenv("TEST_LOAD") == "" {
		t.Skip("load tests are skipped because TEST_LOAD env variable is not set")
	}

	gin.SetMode(gin.TestMode)
	handler := gin.Default()

	monitor := ginmetrics.GetMonitor()
	monitor.SetMetricPath("/metrics")
	monitor.SetSlowTime(1)
	monitor.SetDuration([]float64{0.05, 0.1, 0.25, 0.5, 0.75, 1, 1.5, 2, 3})
	monitor.Use(handler)

	sqlitePath := "./test_load_http.sqlite"
	_ = os.Remove(sqlitePath) // remove to ensure a clean database
	database, err := db.SetupDB(sqlitePath)
	require.NoError(t, err, "setting up db failed")

	service := db.NewDatabase(goqu.New("sqlite3", database))
	server := api.NewServer(service)
	api.SetupRouting(handler, server)

	testServer := httptest.NewServer(handler)

	dr := driver.NewHTTPDriver(testServer.URL, http.DefaultTransport)

	specs.TestLoad(t, dr)
}
