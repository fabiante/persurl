package config

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	path := ".env"

	// check if .env file exists - if not, exit early.
	_, err := os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		return
	}

	err = godotenv.Load(path)
	if err != nil {
		panic(fmt.Errorf("loading env failed: %w", err))
	}
}

func DataDir() string {
	dataDir := os.Getenv(" PERSURL_DATA_DIR")
	if dataDir == "" {
		dataDir = "."
	}
	log.Printf("using data dir: %s", dataDir)
	return dataDir
}

func DbFile(dataDir string) string {
	dbFile := fmt.Sprintf("%s/prod.sqlite", dataDir)
	log.Printf("using database file: %s", dbFile)
	return dbFile
}

func DbDSN() string {
	dsn := os.Getenv("PERSURL_DB_DSN")
	if dsn == "" {
		log.Fatalf("persurl db dsn may not be empty")
	}
	return dsn
}
