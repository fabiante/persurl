package driver

type Driver interface {
	// WithAuth sets the authentication context to be used by this driver.
	WithAuth()
}
