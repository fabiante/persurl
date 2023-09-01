package api

import (
	"github.com/gin-gonic/gin"
)

func SetupRouting(r gin.IRouter, s *Server) {
	r.Use(validPathVar("domain", regexNamed))
	r.Use(validPathVar("name", regexNamed))

	// Resolve endpoints
	r.GET("/r/:domain/:name", s.Resolve)

	// Admin endpoints
	r.PUT("/a/domains/:domain/purls/:name", s.SavePURL)
}
