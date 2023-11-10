package api

import (
	"errors"

	"github.com/fabiante/persurl/api/res"
	"github.com/gin-gonic/gin"
)

var (
	ErrForbidden = errors.New("you are not allowed to do this")
)

// respondWithError responds with an error and aborts the request.
func respondWithError(ctx *gin.Context, status int, err error) {
	response := res.ErrorList{
		Errors: []res.Error{
			err.Error(),
		},
	}

	ctx.AbortWithStatusJSON(status, response)
}
