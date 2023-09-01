package api

import (
	"fmt"
	"regexp"

	"github.com/gin-gonic/gin"
)

var regexNamed *regexp.Regexp

func init() {
	// regexNamed is used to validate everything that has a name. See OpenAPI
	// for more information.
	regexNamed = regexp.MustCompile(`^[a-zA-Z0-9_-]+$`)
}

// validPathVar is a middleware that validates a path variable against a
// regular expression. If the path variable does not match the regular
// expression, the middleware aborts the request with status 400.
func validPathVar(key string, regex *regexp.Regexp) gin.HandlerFunc {
	return func(context *gin.Context) {
		if !regex.MatchString(context.Param(key)) {
			err := fmt.Sprintf("path variable %q does not match regex %s", key, regex.String())
			context.AbortWithStatusJSON(400, err)
			return
		}
		context.Next()
	}
}
