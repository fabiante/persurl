package dsl

import "net/url"

// AdminAPI defines admin features of the application.
type AdminAPI interface {
	// SavePURL creates a new or updates an existing purl.
	SavePURL(purl *PURL) error
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
