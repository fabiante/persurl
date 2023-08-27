package dsl

import (
	"github.com/stretchr/testify/require"
	"testing"
)

// Service defines admin features of the application.
type Service interface {
	CreatePurl(purl *PURL) error
}

// GivenExistingPURL ensures that a PURL is known to the application.
// This is done by simply creating it.
func GivenExistingPURL(t *testing.T, service Service, purl *PURL) {
	err := service.CreatePurl(purl)
	require.NoError(t, err, "creating purl failed")
}
