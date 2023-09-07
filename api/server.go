package api

import (
	"errors"
	"net/http"

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
		ctx.Status(404)
		return
	default:
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
	}
}

func (s *Server) SavePURL(ctx *gin.Context) {
	domain := ctx.Param("domain")
	name := ctx.Param("name")

	type body struct {
		Target string
	}
	var bod body
	if err := ctx.BindJSON(&bod); err != nil {
		panic(err)
	}

	s.service.SavePURL(domain, name, bod.Target)

	ctx.Status(http.StatusNoContent)
}
