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

	row := s.db.QueryRow("SELECT target FROM purls p join domains d on d.id = p.domain_id WHERE d.name = ? AND p.name = ?", domain, name)
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

	// get domain by id
	row := s.db.QueryRow("SELECT id FROM domains WHERE name = ?", domain)
	var domainId int
	err := row.Scan(&domainId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, "domain does not exist")
		} else {
			_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		}
		return
	}

	// insert purl
	_, err = s.db.Exec("INSERT INTO purls (domain_id, name, target) VALUES (?, ?, ?)", domainId, name, bod.Target)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.Status(http.StatusNoContent)
}
