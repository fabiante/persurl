package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouting(r gin.IRouter, s *Server) {
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
		admin.POST("/domains/:domain", s.CreateDomain)

		// PURL
		admin.PUT("/domains/:domain/purls/:name", validName, s.SavePURL)
	}

	// System endpoints
	{
		sys := r.Group("/s")

		sys.GET("/health", func(ctx *gin.Context) {
			// currently no dedicated health check exists.
			// in the future this should be extended with actual checks which signal if the application is ready
			// to receive requests or not.

			ctx.Status(http.StatusNoContent)
		})
	}
}
