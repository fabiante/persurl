package app

import "github.com/fabiante/persurl/app/models"

type ServiceInterface interface {
	AdminServiceInterface
	ResolveServiceInterface
}

type ResolveServiceInterface interface {
	// Resolve tries to resolve a PURL based on the domain and purl name.
	//
	// ErrNotFound is returned if nothing was found.
	Resolve(domain, name string) (string, error)
}

type AdminServiceInterface interface {
	// SavePURL saves a PURL for the given domain name.
	//
	// ErrBadRequest is returned if any parameter is invalid or the domain
	// does not exist.
	SavePURL(domain, name, target string) error

	// CreateDomain creates a new domain.
	//
	// ErrBadRequest is returned if the domain already exists.
	CreateDomain(domain string) error

	// GetDomain returns the domain with the given name.
	//
	// ErrNotFound is returned if the domain does not exist.
	GetDomain(name string) (*models.Domain, error)
}
