package memory_db

import (
	"context"
)

//go:generate ../../bin/mockery --output ./mocks --inpackage-suffix --all --case snake

// Client interface for client
type Client interface {
	DB() DB
	Close() error
}

// DB interface for database
type DB interface {
	QueryExecutor
	ReplyConverter
	Close() error
}

// QueryExecutor interface for requests
type QueryExecutor interface {
	DoContext(ctx context.Context, commandName string, args ...interface{}) (reply interface{}, err error)
}

// ReplyConverter interface for response convert
type ReplyConverter interface {
	String(reply interface{}, err error) (string, error)
}
