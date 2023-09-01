package dsl

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// AdminAPI defines admin features of the application.
type AdminAPI interface {
	CreatePurl(purl *PURL) error
}

// GivenExistingPURL ensures that a PURL is known to the application.
// This is done by simply creating it.
func GivenExistingPURL(t *testing.T, service AdminAPI, purl *PURL) {
	err := service.CreatePurl(purl)
	require.NoError(t, err, "creating purl failed")
}
