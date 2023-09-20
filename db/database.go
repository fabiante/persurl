package db

import (
	"errors"
	"fmt"

	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exec"
	"github.com/fabiante/persurl/app"
	"github.com/lib/pq"
)

// Database implements the applications core logic.
type Database struct {
	db *goqu.Database
}

func NewDatabase(db *goqu.Database) *Database {
	return &Database{
		db: db,
	}
}

func (db *Database) Resolve(domain, name string) (string, error) {
	query := db.db.Select("purls.target").
		From("purls").
		Join(goqu.T("domains"), goqu.On(goqu.I("domains.id").Eq(goqu.I("purls.domain_id")))).
		Where(goqu.And(
			goqu.I("domains.name").Eq(domain),
			goqu.I("purls.name").Eq(name),
		)).
		Limit(1)

	executor := query.Executor()
	var target string
	if found, err := executor.ScanVal(&target); err != nil {
		return "", mapDBError(err)
	} else if !found {
		return "", app.ErrNotFound
	} else {
		return target, nil
	}
}

func (db *Database) SavePURL(domain, name, target string) error {
	// lookup domain first
	query := db.db.Select("id").From("domains").Where(goqu.C("name").Eq(domain)).Limit(1)

	var domainId int
	if found, err := query.ScanVal(&domainId); err != nil {
		return fmt.Errorf("domain lookup faild: %w", err)
	} else if !found {
		return fmt.Errorf("%w: domain does not exist", app.ErrBadRequest)
	}

	// try to get existing purl
	query = db.db.Select("id").From("purls").Where(
		goqu.C("name").Eq(name),
		goqu.C("domain_id").Eq(domainId),
	).Limit(1)

	var upsert interface {
		Executor() exec.QueryExecutor
	}

	var purlId int
	if found, err := query.ScanVal(&purlId); err != nil {
		return fmt.Errorf("purl lookup faild: %w", err)
	} else if !found {
		upsert = db.db.Insert("purls").Rows(goqu.Record{
			"domain_id": domainId,
			"name":      name,
			"target":    target,
		})
	} else {
		upsert = db.db.Update("purls").Set(goqu.Record{
			"target": target,
		}).Where(
			goqu.C("domain_id").Eq(domainId),
			goqu.C("id").Eq(purlId),
		)
	}

	_, err := upsert.Executor().Exec()
	if err != nil {
		return mapDBError(err)
	}

	return nil
}

func (db *Database) CreateDomain(domain string) error {
	stmt := db.db.Insert("domains").
		Cols("name").
		Vals(goqu.Vals{domain})

	if _, err := stmt.Executor().Exec(); err != nil {
		return mapDBError(err)
	} else {
		return nil
	}
}

func (db *Database) DetermineServiceStats() (*app.Stats, error) {
	stat := &app.Stats{}
	var errs []error
	var err error

	stat.DomainsTotal, err = db.countRows("domains", "id")
	errs = append(errs, err)

	stat.PurlsTotal, err = db.countRows("purls", "id")
	errs = append(errs, err)

	return stat, errors.Join(errs...)
}

func (db *Database) countRows(table string, column string) (int, error) {
	var count int

	found, err := db.db.From("domains").Select(goqu.COUNT(goqu.C("id"))).ScanVal(&count)
	switch {
	case err != nil:
		return 0, fmt.Errorf("failed to count rows of table %s: %w", table, err)
	case !found:
		return 0, fmt.Errorf("failed to count rows of table %s, no rows found: %w", table, err)
	}

	return count, nil
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
		return fmt.Errorf("%w: %s", app.ErrBadRequest, err)
	default:
		return fmt.Errorf("unexpected error: %w", err)
	}
}
