package api

import (
	"errors"
	"net/http"

	"github.com/fabiante/persurl/api/res"
	"github.com/fabiante/persurl/app"
	"github.com/gin-gonic/gin"
)

type Server struct {
	service *app.Service
}

func NewServer(service *app.Service) *Server {
	return &Server{service: service}
}

func (s *Server) Resolve(ctx *gin.Context) {
	domain := ctx.Param("domain")
	name := ctx.Param("name")

	target, err := s.service.Resolve(domain, name)
	switch true {
	case err == nil:
		ctx.Redirect(http.StatusFound, target)
		return
	case errors.Is(err, app.ErrNotFound):
		respondWithError(ctx, http.StatusNotFound, err)
		return
	default:
		respondWithError(ctx, http.StatusInternalServerError, err)
	}
}

func (s *Server) SavePURL(ctx *gin.Context) {
	domain := ctx.Param("domain")
	name := ctx.Param("name")

	var req res.SavePURL
	if err := ctx.BindJSON(&req); err != nil {
		ctx.Abort()
		return
	}

	err := s.service.SavePURL(domain, name, req.Target)
	switch true {
	case err == nil:
		ctx.Status(http.StatusNoContent)
	case errors.Is(err, app.ErrBadRequest):
		respondWithError(ctx, http.StatusBadRequest, err)
	default:
		respondWithError(ctx, http.StatusInternalServerError, err)
	}
}

func (s *Server) CreateDomain(ctx *gin.Context) {
	domain := ctx.Param("domain")

	err := s.service.CreateDomain(domain)
	switch true {
	case err == nil:
		ctx.Status(http.StatusNoContent)
	case errors.Is(err, app.ErrBadRequest):
		respondWithError(ctx, http.StatusBadRequest, err)
	default:
		respondWithError(ctx, http.StatusInternalServerError, err)
	}
}
