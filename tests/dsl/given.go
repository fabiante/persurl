package dsl

import (
	"testing"

	"github.com/fabiante/persurl/app"
	"github.com/fabiante/persurl/app/models"
	"github.com/stretchr/testify/require"
)

// GivenExistingPURL ensures that a PURL is known to the application.
//
// This is done by simply creating it.
func GivenExistingPURL(t *testing.T, service AdminAPI, purl *PURL) {
	path, err := service.SavePURL(purl)
	require.NoError(t, err, "saving purl failed")
	require.NotEmpty(t, path)
}

// GivenExistingDomain ensures that a Domain is known to the application.
//
// This currently is a no-op since domains can't explicitly be created.
func GivenExistingDomain(t *testing.T, service AdminAPI, domain string) {
	err := service.CreateDomain(domain)
	require.NoError(t, err, "creating domain failed")
}

// GivenSomeUser creates a user and returns the key for it.
func GivenSomeUser(t *testing.T, userService *app.UserService) *models.UserKey {
	user, err := userService.CreateUser("test@local.com")
	require.NoError(t, err)

	key, err := userService.CreateUserKey(user)
	return key
}
