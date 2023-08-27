package tests

import (
	"github.com/fabiante/persurl/tests/driver"
	"github.com/fabiante/persurl/tests/specs"
	"testing"
)

func TestWithMockDriver(t *testing.T) {
	specs.TestResolver(t, driver.NewMockDriver())
}
