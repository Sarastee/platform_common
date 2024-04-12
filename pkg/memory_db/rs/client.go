package rs

import (
	"github.com/gomodule/redigo/redis"
	"github.com/sarastee/platform_common/pkg/memory_db"
)

var _ memory_db.Client = (*redisClient)(nil)

type redisClient struct {
	masterDB memory_db.DB
}

// New creates new client for redis
func New(pool *redis.Pool) memory_db.Client {
	return &redisClient{masterDB: NewRs(pool)}
}

// DB interface for working with database
func (c *redisClient) DB() memory_db.DB {
	return c.masterDB
}

// Close method closes connection with database
func (c *redisClient) Close() error {
	if c.masterDB != nil {
		return c.masterDB.Close()
	}

	return nil
}
