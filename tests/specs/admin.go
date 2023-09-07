package specs

import (
	"fmt"
	"testing"

	"github.com/fabiante/persurl/app"
	"github.com/fabiante/persurl/tests/dsl"
	"github.com/stretchr/testify/require"
)

func TestAdministration(t *testing.T, admin dsl.AdminAPI) {
	t.Run("administration", func(t *testing.T) {
		t.Run("can't create invalid domain", func(t *testing.T) {
			invalid := []string{
				// empty
				"",
				// whitespace
				"a b",
				// url encoded whitespace
				"a%20b",
				// random characters
				"^",
				"~",
				":",
				",",
				"`",
			}

			for i, domain := range invalid {
				t.Run(fmt.Sprintf("invalid[%d]", i), func(t *testing.T) {
					err := admin.CreateDomain(domain)
					require.Error(t, err)
					require.ErrorIs(t, err, app.ErrBadRequest)
				})
			}
		})

		t.Run("can't create invalid PURL", func(t *testing.T) {
			invalid := []*dsl.PURL{
				// empty
				dsl.NewPURL("valid", "", mustParseURL("example.com")),
				// whitespace
				dsl.NewPURL("valid", "a b", mustParseURL("example.com")),
				// url encoded whitespace
				dsl.NewPURL("valid", "a%20b", mustParseURL("example.com")),
				// random characters
				dsl.NewPURL("valid", "^", mustParseURL("example.com")),
				dsl.NewPURL("valid", "~", mustParseURL("example.com")),
				dsl.NewPURL("valid", ":", mustParseURL("example.com")),
				dsl.NewPURL("valid", ",", mustParseURL("example.com")),
				dsl.NewPURL("valid", "`", mustParseURL("example.com")),
			}

			dsl.GivenExistingDomain(t, admin, "valid")

			for i, purl := range invalid {
				t.Run(fmt.Sprintf("invalid[%d]", i), func(t *testing.T) {
					err := admin.SavePURL(purl)
					require.Error(t, err)
					require.ErrorIs(t, err, app.ErrBadRequest)
				})
			}
		})

		t.Run("can create new PURL", func(t *testing.T) {
			domain := "my-domain-123456"
			purl := dsl.NewPURL(domain, "my-name3456345663456", mustParseURL("https://google.com"))

			dsl.GivenExistingDomain(t, admin, domain)
			// TODO: Assert non-existence of purl to be created
			dsl.GivenExistingPURL(t, admin, purl)
		})

		t.Run("can update existing purl", func(t *testing.T) {
			domain := "my-domain-123456789"
			purl := dsl.NewPURL(domain, "my-name3458904562454564565467", mustParseURL("https://google.com"))

			dsl.GivenExistingDomain(t, admin, domain)
			dsl.GivenExistingPURL(t, admin, purl)

			require.NoError(t, admin.SavePURL(purl), "updating existing purl failed")
		})
	})
}
