package dsl

import (
	"net/url"
)

type PURL struct {
	Domain string
	Name   string
	Target *url.URL
}

func NewPURL(domain string, name string, target *url.URL) *PURL {
	return &PURL{Domain: domain, Name: name, Target: target}
}
