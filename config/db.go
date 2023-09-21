package config

import "errors"

type DB struct {
	DSN            string `mapstructure:"dsn"`
	MaxConnections int    `mapstructure:"max_connections"`
}

// DbDSN is deprecated. Use Get() to retrieve the typed configuration option.
func DbDSN() string {
	dsn := config.DB.DSN

	// TODO: Move this validation into Get() or the init() function. We should ensure that configs are valid.
	if dsn == "" {
		panic(errors.New("db dsn may not be empty"))
	}

	return dsn
}

// DbMaxConnections is deprecated. Use Get() to retrieve the typed configuration option.
func DbMaxConnections() int {
	return config.DB.MaxConnections
}
