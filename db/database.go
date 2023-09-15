package db

import (
	"errors"
	"fmt"
	"sync"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/sqlite3"
	"github.com/doug-martin/goqu/v9/exec"
	"github.com/fabiante/persurl/app"
	"modernc.org/sqlite"
	sqlite3 "modernc.org/sqlite/lib"
)

// Database implements the applications core logic.
type Database struct {
	db   *goqu.Database
	lock *sync.RWMutex
}

func NewDatabase(db *goqu.Database) *Database {
	return &Database{
		db:   db,
		lock: &sync.RWMutex{},
	}
}

func (db *Database) Resolve(domain, name string) (string, error) {
	db.lock.RLock()
	defer db.lock.RUnlock()

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
	db.lock.Lock()
	defer db.lock.Unlock()

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
	db.lock.Lock()
	defer db.lock.Unlock()

	stmt := db.db.Insert("domains").
		Cols("name").
		Vals(goqu.Vals{domain})

	if _, err := stmt.Executor().Exec(); err != nil {
		return mapDBError(err)
	} else {
		return nil
	}
}

func mapDBError(err error) error {
	var serr *sqlite.Error
	if !errors.As(err, &serr) {
		return err
	}

	code := serr.Code()
	switch code {
	case sqlite3.SQLITE_CONSTRAINT_UNIQUE:
		return fmt.Errorf("%w: %s", app.ErrBadRequest, err)
	default:
		return fmt.Errorf("unexpected error: %w", err)
	}
}
