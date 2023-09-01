package specs

import (
	"net/url"
	"testing"

	"github.com/fabiante/persurl/tests/dsl"
	"github.com/stretchr/testify/require"
)

type ResolveAPI interface {
	dsl.AdminAPI

	// ResolvePURL resolves the PURL identified by domain and name, returning
	// the target of the resolved PURL.
	ResolvePURL(domain string, name string) (*url.URL, error)
}

func TestResolver(t *testing.T, resolver ResolveAPI) {
	t.Run("resolver", func(t *testing.T) {
		t.Run("does not resolve non-existant PURL", func(t *testing.T) {
			purl, err := resolver.ResolvePURL("something-very-stupid", "should-not-exist")
			require.Error(t, err)
			require.ErrorIs(t, err, dsl.ErrNotFound)
			require.Nil(t, purl)
		})

		t.Run("resolves existing PURL", func(t *testing.T) {
			domain := "my-domain"
			name := "my-name"

			dsl.GivenExistingPURL(t, resolver, dsl.NewPURL(domain, name, mustParseURL("https://google.com")))

			purl, err := resolver.ResolvePURL(domain, name)
			require.NoError(t, err)
			require.NotNil(t, purl)
		})
	})
}
