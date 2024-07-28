package api

import (
	_ "embed"
	"net/http"

	"github.com/DEXPRO-Solutions-GmbH/swaggerui"
	"github.com/fabiante/persurl/webapp"
	"github.com/gin-gonic/gin"
)

//go:embed openapi.yml
var openAPI []byte

func SetupRouting(r gin.IRouter, s *Server) {
	if swaggerUI, err := swaggerui.NewHandler(openAPI, swaggerui.WithReplaceServerUrls()); err != nil {
		panic(err)
	} else {
		swaggerUI.Register(r)
	}

	isAuthenticated := authenticatedMiddleware(s.user)

	webapp.Register(r)

	validDomain := validPathVar("domain", regexNamed)
	validName := validPathVar("name", regexNamed)

	// Resolve endpoints
	{
		resolve := r.Group("/r")

		resolve.Use(validDomain, validName)

		resolve.GET("/:domain/:name", s.Resolve)
	}

	// Admin endpoints
	{
		admin := r.Group("/a")

		admin.Use(validDomain)

		// Domain
		admin.POST("/domains/:domain", isAuthenticated, s.CreateDomain)

		// PURL
		admin.PUT("/domains/:domain/purls/:name", isAuthenticated, validName, s.SavePURL)
	}

	// System endpoints
	{
		sys := r.Group("/s")

		sys.GET("/health", func(ctx *gin.Context) {
			// currently no dedicated health check exists.
			// in the future this should be extended with actual checks which signal if the application is ready
			// to receive requests or not.

			ctx.JSON(http.StatusOK, "service is ready to receive requests")
		})
	}
}
