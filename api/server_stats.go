package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) GetPublicStats(ctx *gin.Context) {
	stats, err := s.service.DetermineServiceStats()

	switch true {
	case err == nil:
		ctx.JSON(http.StatusOK, stats)
	default:
		respondWithError(ctx, http.StatusInternalServerError, err)
	}
}
