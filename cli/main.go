package main

import (
	"log"

	"github.com/fabiante/persurl/api"
	"github.com/gin-gonic/gin"
)

// main is a crude CLI entrypoint for the application. Will be replaced
// with a proper CLI supporting multiple commands later.
func main() {
	engine := gin.Default()
	server := api.NewServer()
	api.SetupRouting(engine, server)
	if err := engine.Run(":8060"); err != nil {
		log.Fatalf("running api failed: %s", err)
	}
}
