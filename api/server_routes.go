package api

import "github.com/gin-gonic/gin"

func SetupRouting(r gin.IRouter, s *Server) {
	r.GET("/r/:domain/:name", s.Resolve)

	r.PUT("/a/domains/:domain/purls/:name", s.Save)
}
