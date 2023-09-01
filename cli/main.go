package main

import (
	"github.com/fabiante/persurl/api"
	"github.com/gin-gonic/gin"
	"log"
)

// main is a crude CLI entrypoint for the application. Will be replaced
// with a proper CLI supporting multiple commands later.
func main() {
	engine := gin.Default()
	server := api.NewServer()
	api.SetupRouting(engine, server)
	if err := engine.Run(":80"); err != nil {
		log.Fatalf("running api failed: %s", err)
	}
}
