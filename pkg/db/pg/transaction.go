package pg

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/sarastee/platform_common/pkg/db"
)

type manager struct {
	db db.Transactor
}

// NewTransactionManager creates a new transaction manager that satisfies the db.TxManager interface
func NewTransactionManager(db db.Transactor) db.TxManager {
	return &manager{
		db: db,
	}
}

// transaction main function that executes a user-specified handler in a transaction
func (m *manager) transaction(ctx context.Context, opts pgx.TxOptions, fn db.Handler) (err error) {
	// If this is a nested transaction, we skip the initiation of a new transaction and execute the handler.
	tx, ok := ContextTx(ctx)
	if ok {
		return fn(ctx)
	}

	// Starting new transaction
	tx, err = m.db.BeginTx(ctx, opts)
	if err != nil {
		return fmt.Errorf("unable to start transaction: %w", err)
	}

	// Putting transaction in context.
	ctx = MakeContextTx(ctx, tx)

	// We configure the deferment function for rolling back or committing a transaction.
	defer func() {
		// recover after panic
		if r := recover(); r != nil {
			err = fmt.Errorf("panic caught: %v", r)
		}

		// rollback transaction due error
		if err != nil {
			if errRollback := tx.Rollback(ctx); errRollback != nil {
				err = fmt.Errorf("unable to rollback: %w", err)
			}

			return
		}

		// commit transaction, if there are no errors
		if nil == err {
			err = tx.Commit(ctx)
			if err != nil {
				err = fmt.Errorf("unable to commit: %w", err)
			}
		}
	}()

	// Execute code inside a transaction.
	// If the function fails, return an error and the defer function rolls back
	// or otherwise the transaction is committed
	if err = fn(ctx); err != nil {
		err = fmt.Errorf("error while running code inside a transaction: %w", err)
	}

	return err
}

func (m *manager) ReadCommitted(ctx context.Context, f db.Handler) error {
	txOpts := pgx.TxOptions{IsoLevel: pgx.ReadCommitted}
	return m.transaction(ctx, txOpts, f)
}
