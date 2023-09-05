package tests

import (
	"testing"

	"github.com/fabiante/persurl/tests/driver"
	"github.com/fabiante/persurl/tests/specs"
)

func TestWithMockDriver(t *testing.T) {
	d := driver.NewMockDriver()
	specs.TestResolver(t, d)
	specs.TestAdministration(t, d)
}
