package tests

import (
	"testing"

	"github.com/fabiante/persurl/tests/driver"
	"github.com/fabiante/persurl/tests/specs"
)

func TestWithMockDriver(t *testing.T) {
	specs.TestResolver(t, driver.NewMockDriver())
}
