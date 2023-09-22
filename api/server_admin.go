package api

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/fabiante/persurl/api/res"
	"github.com/fabiante/persurl/app"
	"github.com/gin-gonic/gin"
)

func (s *Server) SavePURL(ctx *gin.Context) {
	domain := ctx.Param("domain")
	name := ctx.Param("name")

	var req res.SavePURL
	if err := ctx.BindJSON(&req); err != nil {
		ctx.Abort()
		return
	}

	err := s.admin.SavePURL(domain, name, req.Target)
	switch true {
	case err == nil:
		break
	case errors.Is(err, app.ErrBadRequest):
		respondWithError(ctx, http.StatusBadRequest, err)
	default:
		respondWithError(ctx, http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, res.NewSavePURLResponse(fmt.Sprintf("/r/%s/%s", domain, name)))
}

func (s *Server) CreateDomain(ctx *gin.Context) {
	domain := ctx.Param("domain")

	err := s.admin.CreateDomain(domain)
	switch true {
	case err == nil:
		ctx.Status(http.StatusNoContent)
	case errors.Is(err, app.ErrBadRequest):
		respondWithError(ctx, http.StatusBadRequest, err)
	default:
		respondWithError(ctx, http.StatusInternalServerError, err)
	}
}
