package api

import "github.com/gin-gonic/gin"

func SetupRouting(r gin.IRouter, s *Server) {
	// Resolve endpoints
	r.GET("/r/:domain/:name", s.Resolve)

	// Admin endpoints
	r.PUT("/a/domains/:domain/purls/:name", s.Save)
}
