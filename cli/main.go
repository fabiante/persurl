package main

import (
	"log"

	"github.com/doug-martin/goqu/v9"
	"github.com/fabiante/persurl/api"
	"github.com/fabiante/persurl/db"
	"github.com/gin-gonic/gin"
)

// main is a crude CLI entrypoint for the application. Will be replaced
// with a proper CLI supporting multiple commands later.
func main() {
	database, err := db.SetupDB("./prod.sqlite")
	if err != nil {
		log.Fatalf("setting up database failed: %s", err)
	}

	engine := gin.Default()
	service := db.NewDatabase(goqu.New("sqlite3", database))
	server := api.NewServer(service)
	api.SetupRouting(engine, server)
	if err := engine.Run(":8060"); err != nil {
		log.Fatalf("running api failed: %s", err)
	}
}
