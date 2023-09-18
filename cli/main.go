package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fabiante/persurl/api"
	"github.com/fabiante/persurl/db"
	"github.com/gin-gonic/gin"
)

// main is a crude CLI entrypoint for the application. Will be replaced
// with a proper CLI supporting multiple commands later.
func main() {
	dataDir := envDataDir()

	dbFile := envDbFile(dataDir)

	_, database, err := db.SetupAndMigrateDB(dbFile)
	if err != nil {
		log.Fatalf("setting up database failed: %s", err)
	}

	engine := gin.Default()
	service := db.NewDatabase(database)
	server := api.NewServer(service)
	api.SetupRouting(engine, server)
	if err := engine.Run(":8060"); err != nil {
		log.Fatalf("running api failed: %s", err)
	}
}

func envDataDir() string {
	dataDir := os.Getenv("PERSURL_DATA_DIR")
	if dataDir == "" {
		dataDir = "."
	}
	log.Printf("using data dir: %s", dataDir)
	return dataDir
}

func envDbFile(dataDir string) string {
	dbFile := fmt.Sprintf("%s/prod.sqlite", dataDir)
	log.Printf("using database file: %s", dbFile)
	return dbFile
}
