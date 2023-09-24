package api

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/fabiante/persurl/app"
	"github.com/fabiante/persurl/app/models"
	"github.com/gin-gonic/gin"
)

const (
	headerToken = "Persurl-Key"
)

func authenticatedMiddleware(service app.UserServiceInterface) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader(headerToken)
		if token == "" {
			respondWithError(ctx, http.StatusUnauthorized, errors.New("unauthorized"))
			return
		}

		// load user by email
		user, err := service.GetUserByKey(token)
		switch {
		case errors.Is(err, app.ErrNotFound):
			respondWithError(ctx, http.StatusUnauthorized, errors.New("user not found"))
			return
		case err != nil:
			respondWithError(ctx, http.StatusInternalServerError, err)
			return
		}

		ctx.Set("user", user)

		ctx.Next()
	}
}

func getAuthenticatedUser(ctx *gin.Context) *models.User {
	value, exists := ctx.Get("user")
	if !exists {
		panic(fmt.Errorf("missing context value for key 'user'. ensure you use the appropriate middleware on this endpoint"))
	}

	typed, ok := value.(*models.User)
	if !ok {
		panic(fmt.Errorf("context value has unexpected type %T", value))
	}

	return typed
}
