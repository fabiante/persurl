package app

import (
	"fmt"
)

type serviceMap = map[string]string

// Service is an in-memory implementation of this application's
// features.
//
// It meant as short-term type used until an actual persistence layer is required.
type Service struct {
	data serviceMap
}

func NewService() *Service {
	return &Service{
		data: make(serviceMap),
	}
}

func (s *Service) Resolve(domain, name string) (string, error) {
	target := s.data[fmt.Sprintf("%s/%s", domain, name)]
	if target != "" {
		return target, nil
	}
	return "", ErrNotFound
}

func (s *Service) SavePURL(domain, name string, target string) {
	s.data[fmt.Sprintf("%s/%s", domain, name)] = target
}

func (s *Service) CreateDomain(domain string) error {
	// currently a no-op
	return nil
}
