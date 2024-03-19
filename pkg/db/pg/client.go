package pg

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"github.com/sarastee/platform_common/pkg/db"
)

type pgClient struct {
	masterDBC db.DB
	logger    *zerolog.Logger
}

// New - new Client for Postgres database
func New(ctx context.Context, dsn string, logger *zerolog.Logger) (db.Client, error) {
	pgxConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("error while creating config for pgxpool: %v", err)
	}

	pool, err := pgxpool.NewWithConfig(ctx, pgxConfig)
	if err != nil {
		return nil, fmt.Errorf("error while connecting to DB: %v", err)
	}

	return &pgClient{
		masterDBC: NewDB(pool, logger),
		logger:    logger,
	}, nil
}

// DB access to DB interface
func (c *pgClient) DB() db.DB {
	return c.masterDBC
}

// Close closing connections
func (c *pgClient) Close() error {
	if c.masterDBC != nil {
		c.masterDBC.Close()
	}

	return nil
}
