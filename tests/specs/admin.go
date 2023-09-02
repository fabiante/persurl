package specs

import (
	"fmt"
	"testing"

	"github.com/fabiante/persurl/tests/dsl"
	"github.com/stretchr/testify/require"
)

func TestAdministration(t *testing.T, admin dsl.AdminAPI) {
	t.Run("administration", func(t *testing.T) {
		t.Run("can't create invalid PURL", func(t *testing.T) {
			invalid := []*dsl.PURL{
				// empty
				dsl.NewPURL("", "valid", mustParseURL("example.com")),
				dsl.NewPURL("valid", "", mustParseURL("example.com")),
				// whitespace
				dsl.NewPURL("a b", "valid", mustParseURL("example.com")),
				dsl.NewPURL("valid", "a b", mustParseURL("example.com")),
				// url encoded whitespace
				dsl.NewPURL("a%20b", "valid", mustParseURL("example.com")),
				dsl.NewPURL("valid", "a%20b", mustParseURL("example.com")),
				// random characters
				dsl.NewPURL("^", "valid", mustParseURL("example.com")),
				dsl.NewPURL("~", "valid", mustParseURL("example.com")),
				dsl.NewPURL(":", "valid", mustParseURL("example.com")),
				dsl.NewPURL(",", "valid", mustParseURL("example.com")),
				dsl.NewPURL("`", "valid", mustParseURL("example.com")),
			}

			for i, purl := range invalid {
				t.Run(fmt.Sprintf("invalid[%d]", i), func(t *testing.T) {
					err := admin.CreatePurl(purl)
					require.Error(t, err)
					require.ErrorIs(t, err, dsl.ErrBadRequest)
				})
			}
		})

		t.Run("can create valid PURL", func(t *testing.T) {
			domain := "my-domain"

			dsl.GivenExistingDomain(t, admin, domain)
			dsl.GivenExistingPURL(t, admin, dsl.NewPURL(domain, "my-name", mustParseURL("https://google.com")))
		})
	})
}
