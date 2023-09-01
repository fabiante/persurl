package main

import (
	"database/sql"
	"log"

	"github.com/fabiante/persurl/api"
	"github.com/fabiante/persurl/db"
	"github.com/gin-gonic/gin"
	_ "modernc.org/sqlite"
)

// main is a crude CLI entrypoint for the application. Will be replaced
// with a proper CLI supporting multiple commands later.
func main() {
	engine := gin.Default()

	sqlitePath := "./prod.sqlite"

	database, err := sql.Open("sqlite", sqlitePath)
	if err != nil {
		log.Fatalf("opening sqlite database failed: %s", err)
	}

	err = db.MigrateDb(database)
	if err != nil {
		log.Fatalf("migrating sqlite database failed: %s", err)
	}

	server := api.NewServer(database)
	api.SetupRouting(engine, server)
	if err := engine.Run(":8060"); err != nil {
		log.Fatalf("running api failed: %s", err)
	}
}
