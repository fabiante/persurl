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
