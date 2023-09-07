package specs

import (
	"testing"

	"github.com/fabiante/persurl/app"
	"github.com/fabiante/persurl/tests/dsl"
	"github.com/stretchr/testify/require"
)

func TestResolver(t *testing.T, resolver dsl.API) {
	t.Run("resolver", func(t *testing.T) {
		t.Run("does not resolve non-existant PURL", func(t *testing.T) {
			purl, err := resolver.ResolvePURL("something-very-stupid", "should-not-exist")
			require.Error(t, err)
			require.ErrorIs(t, err, app.ErrNotFound)
			require.Nil(t, purl)
		})

		t.Run("resolves existing PURL", func(t *testing.T) {
			domain := "my-domain"
			name := "my-name"

			dsl.GivenExistingDomain(t, resolver, domain)
			dsl.GivenExistingPURL(t, resolver, dsl.NewPURL(domain, name, mustParseURL("https://google.com")))

			purl, err := resolver.ResolvePURL(domain, name)
			require.NoError(t, err)
			require.NotNil(t, purl)
		})
	})
}
