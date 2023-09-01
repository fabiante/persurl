package tests

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/fabiante/persurl/api"
	"github.com/fabiante/persurl/db"
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

	database, err := sql.Open("sqlite", sqlitePath)
	if err != nil {
		panic(err)
	}

	err = db.MigrateDb(database)
	if err != nil {
		panic(err)
	}

	server := api.NewServer(database)
	api.SetupRouting(handler, server)

	testServer := httptest.NewServer(handler)

	dr := driver.NewHTTPDriver(testServer.URL, http.DefaultTransport)

	specs.TestResolver(t, dr)
	specs.TestAdministration(t, dr)
}
