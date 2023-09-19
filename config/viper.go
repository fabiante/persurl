package config

import (
	"github.com/spf13/viper"
)

var vip *viper.Viper

func init() {
	vip = setupViper()
}

func setupViper() *viper.Viper {
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

	// trigger config parsing
	check(v.ReadInConfig())

	return v
}
