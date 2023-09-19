package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

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

func DbDSN() string {
	dsn := os.Getenv("PERSURL_DB_DSN")
	if dsn == "" {
		dsn = os.Getenv("DATABASE_URL")
	}
	if dsn == "" {
		log.Fatalf("persurl db dsn may not be empty")
	}
	return dsn
}

func DbMaxConnections() int {
	val := os.Getenv("PERSURL_DB_MAX_CONNECTIONS")
	if val == "" {
		val = "10"
	}

	maxCon, err := strconv.ParseInt(val, 10, 32)
	if err != nil {
		log.Fatalf("invalid db max connection parameter %s", val)
	}

	return int(maxCon)
}
