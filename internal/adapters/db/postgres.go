package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/spf13/viper"
)

func NewPostgresDB(ctx context.Context, cfg ConnectionConfig) (*pgxpool.Pool, error) {

	connString := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s", cfg.Username, cfg.Password, cfg.Host,
		cfg.Port, cfg.DBName, cfg.SSLMode)

	conf, err := pgxpool.ParseConfig(connString) // Using environment variables instead of a connection string.
	if err != nil {
		return nil, err
	}

	var logLevel pgx.LogLevel
	switch viper.GetString("app.loglevel") {
	case "panic":
		logLevel = pgx.LogLevelError
		break
	case "fatal":
		logLevel = pgx.LogLevelError
		break
	case "error":
		logLevel = pgx.LogLevelError
		break
	case "warning":
		logLevel = pgx.LogLevelWarn
		break
	case "info":
		logLevel = pgx.LogLevelInfo
		break
	case "trace":
		logLevel = pgx.LogLevelTrace
		break
	default:
		logLevel = pgx.LogLevelInfo
		break
	}

	conf.ConnConfig.LogLevel = logLevel
	conf.MaxConns = 50
	conf.ConnConfig.PreferSimpleProtocol = true

	pool, err := pgxpool.ConnectConfig(ctx, conf)
	if err != nil {
		return nil, fmt.Errorf("pgx connection error: %w", err)
	}

	if err = pingDB(ctx, pool); err != nil {
		return nil, err
	}

	return pool, nil
}

func pingDB(ctx context.Context, pool *pgxpool.Pool) error {
	conn, err := pool.Acquire(ctx)
	defer conn.Release()
	if err != nil {
		return err
	}
	if err = conn.Ping(ctx); err != nil {
		return err
	}
	return nil
}
