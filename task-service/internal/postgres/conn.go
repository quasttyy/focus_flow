package postgres

import (
	"context"
	"focus_flow/common/logger"
	"os"
	"time"
	"errors"
	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectPostgres() (*pgxpool.Pool, error) {
    dsn := os.Getenv("DB_URL")
    if dsn == "" {
        return nil, errors.New("DB_URL not found")
    }

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    pool, err := pgxpool.New(ctx, dsn)
    if err != nil {
        logger.Log.Error("failed to create pool:", err)
        return nil, err
    }

    if err := pool.Ping(ctx); err != nil {
        logger.Log.Error("failed to ping postgres:", err)
        return nil, err
    }

    logger.Log.Info("successfully connected to PostgreSQL")
    return pool, nil
}
