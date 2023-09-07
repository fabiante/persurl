package app

import (
	"fmt"
)

type DomainMap map[string]PurlMap

type PurlMap map[string]string

func (m PurlMap) CreatePurl(name string, target string) {
	m[name] = target
}

func (m PurlMap) ResolvePURL(name string) (string, error) {
	if purl, purlExists := m[name]; purlExists {
		return purl, nil
	}

	return "", fmt.Errorf("%w: purl does not exist", ErrNotFound)
}
