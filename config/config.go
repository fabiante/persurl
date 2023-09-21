package config

import "errors"

// init initializes this packages state.
func init() {
	errs := []error{
		initViper(),
		initConfig(),
	}

	if err := errors.Join(errs...); err != nil {
		panic(err)
	}
}

// Config is the top level aggregator for all configurations.
type Config struct {
	DB *DB
}

var config *Config

func initConfig() error {
	config = &Config{
		DB: &DB{},
	}

	errs := []error{
		vip.UnmarshalKey("db", config.DB),
	}

	return errors.Join(errs...)
}

// Get retrieves the singleton Config of this application.
//
// You may store references to this config, but you shall not modify it.
func Get() *Config {
	return config
}
