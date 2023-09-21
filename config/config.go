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

	// TestLoad is used during development to enable load testing.
	TestLoad bool `mapstructure:"test_load"`
}

var config *Config

func initConfig() error {
	config = &Config{
		DB: &DB{},
	}

	errs := []error{
		vip.Unmarshal(config),
	}

	return errors.Join(errs...)
}

// Get retrieves the singleton Config of this application.
//
// You may store references to this config, but you shall not modify it.
func Get() *Config {
	return config
}
