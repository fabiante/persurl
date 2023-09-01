package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	data map[string]string
}

func NewServer() *Server {
	return &Server{
		data: make(map[string]string),
	}
}

func (s *Server) Resolve(ctx *gin.Context) {
	domain := ctx.Param("domain")
	name := ctx.Param("name")
	x := s.data[fmt.Sprintf("%s/%s", domain, name)]
	if x == "" {
		ctx.Status(404)
		return
	} else {
		ctx.Redirect(http.StatusFound, x)
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
	s.data[fmt.Sprintf("%s/%s", domain, name)] = bod.Target
	ctx.Status(http.StatusNoContent)
}
