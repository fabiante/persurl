package config

import (
	"errors"
)

func DbDSN() string {
	dsn := vip.GetString("db.dsn")

	if dsn == "" {
		panic(errors.New("db dsn may not be empty"))
	}

	return dsn
}

func DbMaxConnections() int {
	return vip.GetInt("db.max_connections")
}

func TestLoad() bool {
	return vip.IsSet("test_load")
}
