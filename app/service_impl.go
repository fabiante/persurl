package app

import (
	"errors"
	"fmt"

	"github.com/fabiante/persurl/app/models"
	"gorm.io/gorm"
)

// service implements ServiceInterface based on a SQL database, acessed via gorm.DB
//
// The given gorm.DB instance is expected to have enabled error translation. This is required for properly
// mapping errors onto application specific errors.
type service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) ServiceInterface {
	return &service{db: db}
}

func (s *service) Resolve(domain, name string) (string, error) {
	purl := &models.PURL{}

	err := s.db.Model(&models.PURL{}).
		Joins("join domains on domains.id = purls.domain_id").
		Where("domains.name = ?", domain).
		Where("purls.name = ?", name).
		Take(purl).Error

	switch {
	case err == nil:
		return purl.Target, nil
	case errors.Is(err, gorm.ErrRecordNotFound):
		return "", ErrNotFound
	default:
		return "", mapDBError(err)
	}
}

func (s *service) CreateDomain(name string) error {
	domain := &models.Domain{
		Name: name,
	}

	err := s.db.Create(domain).Error

	if err != nil {
		return mapDBError(err)
	}

	return nil
}

func (s *service) SavePURL(domainName, name, target string) error {
	domain := &models.Domain{}

	// get domain
	{
		err := s.db.Where(&models.Domain{Name: domainName}).Take(domain).Error
		if err != nil {
			switch {
			case err == nil:
				break
			case errors.Is(err, gorm.ErrRecordNotFound):
				return fmt.Errorf("%w: domain does not exist", ErrBadRequest)
			default:
				return mapDBError(err)
			}
		}
	}

	// save purl
	{
		purl := &models.PURL{
			DomainID: domain.ID,
			Name:     name,
			Target:   target,
		}

		err := s.db.FirstOrCreate(purl).Error

		if err != nil {
			return mapDBError(err)
		}
	}

	return nil
}

func mapDBError(err error) error {
	switch {
	case errors.Is(err, gorm.ErrDuplicatedKey):
		return fmt.Errorf("%w: %s", ErrBadRequest, err)
	default:
		return err
	}
}
