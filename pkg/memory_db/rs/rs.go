package rs

import (
	"context"
	"fmt"

	"github.com/gomodule/redigo/redis"
	"github.com/sarastee/platform_common/pkg/memory_db"
)

var _ memory_db.DB = (*rs)(nil)

type rs struct {
	pool *redis.Pool
}

// NewRs interface for database execution
func NewRs(pool *redis.Pool) memory_db.DB {
	return &rs{
		pool: pool,
	}
}

// DoContext executes command with context
func (r *rs) DoContext(ctx context.Context, commandName string, args ...interface{}) (reply interface{}, err error) {
	conn, err := r.pool.GetContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("failure while connecting to Redis from connection pool: %w", err)
	}
	defer func() {
		_ = conn.Close()
	}()

	reply, err = conn.Do(commandName, args...)
	if err != nil {
		return nil, fmt.Errorf("failure while executing command to Redis: %w", err)
	}

	return reply, nil
}

// Close method closes connection with database
func (r *rs) Close() error {
	return r.pool.Close()
}

// String method converts response from database to string
func (r *rs) String(reply interface{}, err error) (string, error) {
	return redis.String(reply, err)
}
