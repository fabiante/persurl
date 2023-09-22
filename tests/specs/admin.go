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
		testDomainAdmin(t, admin)
		testPurlAdmin(t, admin)
	})
}

func testPurlAdmin(t *testing.T, admin dsl.AdminAPI) {
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
				_, err := admin.SavePURL(purl)
				require.Error(t, err)
				//require.ErrorIs(t, err, app.ErrBadRequest) // TODO: Some tests cause a 404 with the http driver.
			})
		}
	})

	t.Run("can't create PURL on non-existent domain", func(t *testing.T) {
		domain := "this-domain-does-not-exist-it-should-not"
		purl := dsl.NewPURL(domain, "my-name3456334654645663456", mustParseURL("https://google.com"))

		_, err := admin.SavePURL(purl)
		require.ErrorIs(t, err, app.ErrBadRequest)
	})

	t.Run("can create new PURL", func(t *testing.T) {
		domain := "my-domain-123456"
		dsl.GivenExistingDomain(t, admin, domain)

		validPurls := []*dsl.PURL{
			dsl.NewPURL(domain, "my-name3456345663456", mustParseURL("https://google.com")),
			dsl.NewPURL(domain, "my-purl.com", mustParseURL("https://google.com")),
		}

		for i, purl := range validPurls {
			t.Run(fmt.Sprintf("valid[%d]", i), func(t *testing.T) {
				// TODO: Assert non-existence of purl to be created
				path, err := admin.SavePURL(purl)
				require.NoError(t, err, "creating purl failed")
				require.NotEmpty(t, path)
			})
		}
	})

	t.Run("can update existing purl", func(t *testing.T) {
		domain := "my-domain-123456789"
		purl := dsl.NewPURL(domain, "my-name3458904562454564565467", mustParseURL("https://google.com"))

		dsl.GivenExistingDomain(t, admin, domain)
		dsl.GivenExistingPURL(t, admin, purl)

		// modify purl's name - updating the target would be the usual case but that is harder to assert.
		purl.Name = "my-new-name-updated"

		path, err := admin.SavePURL(purl)
		require.NoError(t, err, "updating existing purl failed")
		require.NotEmpty(t, path)
		require.Contains(t, path, "my-new-name-updated")
	})
}

func testDomainAdmin(t *testing.T, admin dsl.AdminAPI) {
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
				//require.ErrorIs(t, err, app.ErrBadRequest) // TODO: Some tests cause a 404 with the http driver.
			})
		}
	})

	t.Run("can create valid domain", func(t *testing.T) {
		valid := []string{
			"awesome-domain-unique-name-123",
			"awesome.com",
		}

		for i, v := range valid {
			t.Run(fmt.Sprintf("valid[%d]", i), func(*testing.T) {
				err := admin.CreateDomain(v)
				require.NoError(t, err)
			})
		}
	})

	t.Run("can't create duplicate domain", func(t *testing.T) {
		domain := "should-exist-once-4357824758wr47895645"
		dsl.GivenExistingDomain(t, admin, domain)
		err := admin.CreateDomain("should-exist-once-4357824758wr47895645")
		require.Error(t, err)
		require.ErrorIs(t, err, app.ErrBadRequest)
	})
}
