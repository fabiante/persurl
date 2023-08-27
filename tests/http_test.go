package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fabiante/persurl/api"
	"github.com/fabiante/persurl/tests/driver"
	"github.com/fabiante/persurl/tests/specs"
	"github.com/gin-gonic/gin"
)

func TestWithHTTPDriver(t *testing.T) {
	gin.SetMode(gin.TestMode)
	handler := gin.Default()

	server := api.NewServer()
	api.SetupRouting(handler, server)

	testServer := httptest.NewServer(handler)
	specs.TestResolver(t, driver.NewHTTPDriver(testServer.URL, http.DefaultClient))
}
