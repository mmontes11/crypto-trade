package database

import (
	ctx "context"
	"database/sql"
)

// TxFn function that runs in the transaction
type TxFn = func(ctx ctx.Context, tx *sql.Tx) error

// Tx executes txFn inside a transaction
func Tx(ctx ctx.Context, db *sql.DB, txFn TxFn) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil
	}

	if err := txFn(ctx, tx); err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	return tx.Commit()
}
