package api

import (
	"github.com/fabiante/persurl/api/res"
	"github.com/gin-gonic/gin"
)

func respondWithError(ctx *gin.Context, status int, err error) {
	response := res.ErrorList{
		Errors: []res.Error{
			err.Error(),
		},
	}

	ctx.AbortWithStatusJSON(status, response)
}

type ErrHandler func(c *gin.Context) error

func AsErrHandler(h ErrHandler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := h(ctx)

		if err != nil {
			_ = ctx.AbortWithError(500, err)
			return
		}
	}
}
