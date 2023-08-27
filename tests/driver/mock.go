package driver

import (
	"fmt"
	"github.com/fabiante/persurl/tests/dsl"
	"net/url"
)

// MockDriver is an in-process implementation of a test driver.
//
// The purpose of this type is to quickly write test specifications
// without the overhead an HTTP API.
type MockDriver struct {
	purls map[string]*dsl.PURL
}

func NewMockDriver() *MockDriver {
	return &MockDriver{
		make(map[string]*dsl.PURL),
	}
}

func (m *MockDriver) CreatePurl(purl *dsl.PURL) error {
	m.purls[fmt.Sprintf("%s/%s", purl.Domain, purl.Name)] = purl
	return nil
}

func (m *MockDriver) ResolvePURL(domain string, name string) (*url.URL, error) {
	if purl, found := m.purls[fmt.Sprintf("%s/%s", domain, name)]; found {
		return purl.Target, nil
	} else {
		return nil, dsl.ErrNotFound
	}
}
