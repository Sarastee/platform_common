package db

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

//go:generate ../../bin/mockery --output ./mocks  --inpackage-suffix --all

// Handler - function, that executed in transaction
type Handler func(ctx context.Context) error

// Client - client for db
type Client interface {
	DB() DB
	Close() error
}

// TxManager transaction manager
type TxManager interface {
	ReadCommitted(ctx context.Context, f Handler) error
}

// Query wrapper over request, containing Name and QueryRaw
// Name can be used for logging and tracing
type Query struct {
	Name     string
	QueryRaw string
}

// Transactor interface for transaction
type Transactor interface {
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
}

// SQLExecutor combines NamedExecutor and QueryExecutor
type SQLExecutor interface {
	QueryExecutor
	CopyExecutor
}

// QueryExecutor interface for usual requests
type QueryExecutor interface {
	ExecContext(ctx context.Context, q Query, args ...interface{}) (pgconn.CommandTag, error)
	QueryContext(ctx context.Context, q Query, args ...interface{}) (pgx.Rows, error)
	QueryRowContext(ctx context.Context, q Query, args ...interface{}) pgx.Row
}

// CopyExecutor interface for copying requests
type CopyExecutor interface {
	CopyFromContext(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error)
}

// Pinger interface for pinging database
type Pinger interface {
	Ping(ctx context.Context) error
}

// DB interface for database
type DB interface {
	SQLExecutor
	Transactor
	Pinger
	Close()
}
