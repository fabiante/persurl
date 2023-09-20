package api

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

var regexNamed *regexp.Regexp

func init() {
	// regexNamed is used to validate everything that has a name. See OpenAPI
	// for more information.
	regexNamed = regexp.MustCompile(`^[a-zA-Z0-9\\._-]+$`)
}

// validPathVar is a middleware that validates a path variable against a
// regular expression. If the path variable does not match the regular
// expression, the middleware aborts the request with status 400.
func validPathVar(key string, regex *regexp.Regexp) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !regex.MatchString(ctx.Param(key)) {
			respondWithError(ctx, http.StatusBadRequest, fmt.Errorf("path variable %q does not match regex %s", key, regex.String()))
			return
		}
		ctx.Next()
	}
}
