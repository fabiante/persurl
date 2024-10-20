package api

import (
	"errors"
	"net/http"

	"github.com/fabiante/persurl/app"
	"github.com/gin-gonic/gin"
)

type Server struct {
	resolver app.ResolveServiceInterface
	admin    app.AdminServiceInterface
	user     app.UserServiceInterface
}

func NewServer(resolver app.ResolveServiceInterface, admin app.AdminServiceInterface, user app.UserServiceInterface) *Server {
	return &Server{resolver: resolver, admin: admin, user: user}
}

func (s *Server) Resolve(ctx *gin.Context) error {
	domain := ctx.Param("domain")
	name := ctx.Param("name")

	target, err := s.resolver.Resolve(domain, name)
	switch true {
	case err == nil:
		ctx.Redirect(http.StatusFound, target)
		return nil
	case errors.Is(err, app.ErrNotFound):
		respondWithError(ctx, http.StatusNotFound, err)
		return nil
	default:
		return err
	}
}
