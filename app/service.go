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
	SavePURL(domain *models.Domain, name, target string) error

	// CreateDomain creates a new domain.
	//
	// ErrBadRequest is returned if the domain already exists.
	CreateDomain(domain string) (*models.Domain, error)

	// GetDomain returns the domain with the given name.
	//
	// ErrNotFound is returned if the domain does not exist.
	GetDomain(name string) (*models.Domain, error)
}

type UserServiceInterface interface {
	CreateUser(email string) (*models.User, error)
	GetUser(email string) (*models.User, error)
	GetUserByKey(key string) (*models.User, error)

	CreateUserKey(user *models.User) (*models.UserKey, error)
	GetUserKey(value string) (*models.UserKey, error)
}
