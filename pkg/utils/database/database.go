package database

import (
	"fmt"

	"github.com/efectn/library-management/pkg/database/ent"
)

func EntRollback(tx *ent.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%w: %v", err, rerr)
	}
	return err
}
