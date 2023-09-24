package specs

import (
	"context"
	"testing"

	"github.com/fabiante/persurl/app"
	"github.com/fabiante/persurl/tests/dsl"
	"github.com/stretchr/testify/require"
)

func TestResolver(t *testing.T, resolver dsl.API) {
	ctx := context.TODO()

	t.Run("resolver", func(t *testing.T) {
		t.Run("does not resolve non-existent domain", func(t *testing.T) {
			domain := "something-very-stupid-9873214356"
			name := "should-not-exist"

			purl, err := resolver.ResolvePURL(ctx, domain, name)
			require.Error(t, err)
			require.ErrorIs(t, err, app.ErrNotFound)
			require.Nil(t, purl)
		})

		t.Run("does not resolve non-existent purl", func(t *testing.T) {
			domain := "something-very-stupid-34563456"
			name := "should-not-exist"

			dsl.GivenExistingDomain(ctx, t, resolver, domain)

			purl, err := resolver.ResolvePURL(ctx, domain, name)
			require.Error(t, err)
			require.ErrorIs(t, err, app.ErrNotFound)
			require.Nil(t, purl)
		})

		t.Run("resolves existing PURL", func(t *testing.T) {
			domain := "my-domain"
			name := "my-name"

			dsl.GivenExistingDomain(ctx, t, resolver, domain)
			dsl.GivenExistingPURL(ctx, t, resolver, dsl.NewPURL(domain, name, mustParseURL("https://google.com")))

			purl, err := resolver.ResolvePURL(ctx, domain, name)
			require.NoError(t, err)
			require.NotNil(t, purl)
		})
	})
}
