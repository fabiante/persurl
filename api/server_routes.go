package api

import (
	"github.com/gin-gonic/gin"
)

func SetupRouting(r gin.IRouter, s *Server) {
	r.Use(errHandler)

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
}
