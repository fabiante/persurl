package app

import (
	"errors"
	"fmt"

	"github.com/fabiante/persurl/app/models"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type AdminService struct {
	db *gorm.DB
}

func NewAdminService(db *gorm.DB) *AdminService {
	return &AdminService{db: db}
}

func (s *AdminService) Resolve(domain, name string) (string, error) {
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

func (s *AdminService) CreateDomain(name string) error {
	domain := &models.Domain{
		Name: name,
	}

	err := s.db.Create(domain).Error

	if err != nil {
		return mapDBError(err)
	}

	return nil
}

func (s *AdminService) SavePURL(domainName, name, target string) error {
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

const (
	pgErrUniqueKeyViolation = "23505"
)

func mapDBError(err error) error {
	var serr *pq.Error
	if !errors.As(err, &serr) {
		return err
	}

	// Error codes
	// SQLite: https://www.sqlite.org/rescode.html
	// Postgres: http://www.postgresql.org/docs/9.3/static/errcodes-appendix.html

	code := serr.Code
	switch code {
	case pgErrUniqueKeyViolation:
		return fmt.Errorf("%w: %s", ErrBadRequest, err)
	default:
		return fmt.Errorf("unexpected error: %w", err)
	}
}
