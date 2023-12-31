package config

import (
	"errors"

	"github.com/spf13/viper"
)

var vip *viper.Viper

func initViper() error {
	vip = newViper()

	// trigger config parsing - optional
	if err := vip.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if !errors.As(err, &configFileNotFoundError) {
			return err
		}
	}

	return nil
}

func newViper() *viper.Viper {
	v := viper.New()

	// loading
	v.AddConfigPath(".")
	v.SetConfigName("app")

	// helper to panic on any error
	check := func(e error) {
		if e != nil {
			panic(e)
		}
	}

	// defaults
	v.SetDefault("db.max_connections", 10)

	// env binding
	check(setupEnv(v))

	return v
}
