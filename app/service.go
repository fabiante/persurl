package app

type ServiceInterface interface {
	// Resolve tries to resolve a PURL based on the domain and purl name.
	//
	// ErrNotFound is returned if nothing was found.
	Resolve(domain, name string) (string, error)

	// SavePURL saves a PURL for the given domain name.
	//
	// ErrBadRequest is returned if any parameter is invalid or the domain
	// does not exist.
	SavePURL(domain, name, target string) error

	// CreateDomain creates a new domain.
	//
	// ErrBadRequest is returned if the domain already exists.
	CreateDomain(domain string) error
}
