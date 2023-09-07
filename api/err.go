package api

import (
	"net/http"

	"github.com/fabiante/persurl/app"
	"github.com/gin-gonic/gin"
	err "github.com/richzw/gin-error"
)

// errHandler is a error handling middleware which maps errors onto
// http status codes.
var errHandler gin.HandlerFunc

func init() {
	maps := []*err.ErrorMap{
		err.NewErrMap(app.ErrBadRequest).StatusCode(http.StatusBadRequest),
		err.NewErrMap(app.ErrNotFound).StatusCode(http.StatusNotFound),
	}

	errHandler = err.Error(maps...)
}
