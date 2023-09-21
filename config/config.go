package config

// init initializes this packages state.
func init() {
	if err := initViper(); err != nil {
		panic(err)
	}
}
