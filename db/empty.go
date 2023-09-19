package db

import (
	"errors"

	"github.com/doug-martin/goqu/v9"
)

// EmptyTables is used to empty a collection of tables. This may be useful if truncating a
// table is not possible.
func EmptyTables(db *goqu.Database, tables ...string) error {
	return db.WithTx(func(db *goqu.TxDatabase) error {
		var errs []error
		for _, table := range tables {
			_, err := db.Delete(table).Executor().Exec()
			errs = append(errs, err)
		}
		return errors.Join(errs...)
	})
}
