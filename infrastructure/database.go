package infrastructure

import (
	"context"
	"database/sql"
)

type Database interface {
	Get(ctx context.Context, key string) (string, error)
}

type database struct {
	db *sql.DB
}
