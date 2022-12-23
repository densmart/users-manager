package db

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

type ConnectionConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
	TimeZone string
}

type WrapperDB struct {
	DBType string
	Pool   *pgxpool.Pool
	Ctx    context.Context
}

func NewDB(ctx context.Context, DBType string, cfg ConnectionConfig) (*WrapperDB, error) {
	var pool *pgxpool.Pool
	var err error
	switch DBType {
	case "postgres":
		pool, err = NewPostgresDB(ctx, cfg)
	default:
		pool, err = NewPostgresDB(ctx, cfg)
	}
	if err != nil {
		return nil, err
	}

	return &WrapperDB{
		DBType: DBType,
		Pool:   pool,
		Ctx:    ctx,
	}, nil
}

func (db *WrapperDB) Close() {
	switch db.DBType {
	case "postgres":
		db.Pool.Close()
	default:
		db.Pool.Close()
	}
}
