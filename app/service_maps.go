package app

import (
	"fmt"
	"net/url"

	"github.com/fabiante/persurl/tests/dsl"
)

type DomainMap map[string]PurlMap

type PurlMap map[string]*dsl.PURL

func (m PurlMap) CreatePurl(purl *dsl.PURL) error {
	m[purl.Name] = purl
	return nil
}

func (m PurlMap) ResolvePURL(name string) (*url.URL, error) {
	if purl, purlExists := m[name]; purlExists {
		return purl.Target, nil
	}

	return nil, fmt.Errorf("%w: purl does not exist", ErrNotFound)
}
