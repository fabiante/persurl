package config

import (
	"errors"

	"github.com/spf13/viper"
)

func setupEnv(v *viper.Viper) error {
	errs := []error{
		v.BindEnv("test_load", "TEST_LOAD"),
		v.BindEnv("db.dsn", "PERSURL_DB_DSN", "DATABASE_URL"),
		v.BindEnv("db.max_connections", "PERSURL_DB_MAX_CONNECTIONS"),
	}

	return errors.Join(errs...)
}
