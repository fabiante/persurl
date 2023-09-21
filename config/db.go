package config

import "errors"

type DB struct {
	DSN            string `mapstructure:"dsn"`
	MaxConnections int    `mapstructure:"max_connections"`
}

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
