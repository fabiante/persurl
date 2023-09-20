package webapp

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:embed dist/*
var dist embed.FS

func Register(r gin.IRoutes) {
	r.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusPermanentRedirect, "/webapp")
	})

	subFs, err := fs.Sub(dist, "dist")
	if err != nil {
		panic(err)
	}

	r.StaticFS("/webapp", http.FS(subFs))
}
