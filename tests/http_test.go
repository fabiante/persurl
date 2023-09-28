package tests

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fabiante/persurl/api"
	"github.com/fabiante/persurl/app"
	"github.com/fabiante/persurl/config"
	"github.com/fabiante/persurl/db"
	"github.com/fabiante/persurl/tests/driver"
	"github.com/fabiante/persurl/tests/dsl"
	"github.com/fabiante/persurl/tests/specs"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestWithHTTPDriver(t *testing.T) {
	ctx := context.TODO()
	gin.SetMode(gin.TestMode)
	handler := gin.Default()

	database, err := db.SetupAndMigratePostgresDB(config.DbDSN(), config.DbMaxConnections())
	require.NoError(t, err, "setting up db failed")

	err = db.EmptyTables(database.Goqu, "purls", "domains", "user_keys", "users")
	require.NoError(t, err, "truncating tables failed")

	service := app.NewService(database.Gorm)
	userService := app.NewUserService(database.Gorm)

	key := dsl.GivenSomeUser(ctx, t, userService)
	ctx = driver.CtxWithUserKey(ctx, key.Value)

	server := api.NewServer(service, service, userService)
	api.SetupRouting(handler, server)

	testServer := httptest.NewServer(handler)

	dr := driver.NewHTTPDriver(testServer.URL, http.DefaultTransport)

	specs.TestResolver(t, dr)
	specs.TestAdministration(t, dr)
}
