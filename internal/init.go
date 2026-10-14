package internal

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"

	"booking-api/utils"
)

// Config is filled by Init and readable anywhere: internal.Config.Port
var Config utils.Config

// Init runs once at startup: reads .env, then connects Postgres and Redis.
// Call it first thing in main.
func Init(ctx context.Context) error {
	if err := loadEnv(); err != nil {
		return err
	}
	if err := PgConn(ctx, Config.DatabaseURL); err != nil {
		return err
	}
	if err := RedisConn(ctx, Config.RedisURL); err != nil {
		return err
	}
	return nil
}

// Close shuts down all connections. Call it with defer in main.
func Close() {
	if Redis != nil {
		Redis.Close()
	}
	if DB != nil {
		DB.Close()
	}
}

func loadEnv() error {
	// In production there is no .env file, only real environment variables,
	// so a missing file is not an error.
	_ = godotenv.Load()

	Config = utils.Config{
		Port:        os.Getenv("PORT"),
		DatabaseURL: os.Getenv("DATABASE_URL"),
		RedisURL:    os.Getenv("REDIS_URL"),
	}

	if Config.DatabaseURL == "" {
		return fmt.Errorf("DATABASE_URL is required")
	}
	return nil
}
