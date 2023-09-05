package dsl

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// GivenExistingPURL ensures that a PURL is known to the application.
//
// This is done by simply creating it.
func GivenExistingPURL(t *testing.T, service AdminAPI, purl *PURL) {
	err := service.CreatePurl(purl)
	require.NoError(t, err, "creating purl failed")
}

// GivenExistingDomain ensures that a Domain is known to the application.
//
// This currently is a no-op since domains can't explicitly be created.
func GivenExistingDomain(t *testing.T, service AdminAPI, domain string) {
	// no-op - Implement something here once domains can be created.
}
