package dsl

import "net/url"

// AdminAPI defines admin features of the application.
type AdminAPI interface {
	CreateDomain(name string) error

	// SavePURL creates a new or updates an existing purl.
	//
	// If no error occurred the returned string is the path (without host) of the created PURL.
	SavePURL(purl *PURL) (string, error)
}

// ResolveAPI defines features for PURL resolution.
type ResolveAPI interface {
	// ResolvePURL resolves the PURL identified by domain and name, returning
	// the target of the resolved PURL.
	ResolvePURL(domain string, name string) (*url.URL, error)
}

// API aggregates all apis.
type API interface {
	AdminAPI
	ResolveAPI
}
