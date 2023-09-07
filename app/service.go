package app

import (
	"fmt"
)

// Service is an in-memory implementation of this application's
// features.
//
// It meant as short-term type used until an actual persistence layer is required.
type Service struct {
	data DomainMap
}

func NewService() *Service {
	return &Service{
		data: make(DomainMap),
	}
}

func (s *Service) Resolve(domain, name string) (string, error) {
	purls, found := s.data[domain]
	if !found {
		return "", ErrNotFound
	}
	return purls.ResolvePURL(name)
}

func (s *Service) SavePURL(domain, name, target string) error {
	purls, found := s.data[domain]
	if !found {
		return fmt.Errorf("%w: domain does not exist", ErrBadRequest)
	}

	purls.CreatePurl(name, target)
	return nil
}

func (s *Service) CreateDomain(domain string) error {
	if _, found := s.data[domain]; found {
		return fmt.Errorf("%w: domain already exists", ErrBadRequest)
	}

	s.data[domain] = make(PurlMap)
	return nil
}
