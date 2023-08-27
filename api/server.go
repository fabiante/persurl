package api

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Resolve(ctx *gin.Context) {
	panic("not implemented") // todo
}

func (s *Server) Save(ctx *gin.Context) {
	panic("not implemented") // todo
}
