package api

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	db *sql.DB
}

func NewServer(db *sql.DB) *Server {
	return &Server{
		db: db,
	}
}

func (s *Server) Resolve(ctx *gin.Context) {
	domain := ctx.Param("domain")
	name := ctx.Param("name")

	row := s.db.QueryRow("SELECT target FROM purls WHERE domain = ? AND name = ?", domain, name)
	var target string
	err := row.Scan(&target)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.AbortWithStatus(http.StatusNotFound)
		} else {
			_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		}

		return
	}

	ctx.Redirect(http.StatusFound, target)
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

	_, err := s.db.Exec("INSERT INTO purls (domain, name, target) VALUES (?, ?, ?)", domain, name, bod.Target)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.Status(http.StatusNoContent)
}
