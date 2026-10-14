package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v3"

	"booking-api/internal"
)

func main() {
	// ctx is cancelled when you press Ctrl+C (or the host sends SIGTERM).
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	if err := internal.Init(ctx); err != nil {
		slog.Error("startup failed", "error", err)
		os.Exit(1)
	}
	defer internal.Close()
	slog.Info("connected to postgres and redis")

	app := internal.NewApp()

	// Listen blocks until ctx is cancelled, then gives in-flight requests
	// up to 10 seconds to finish before returning.
	err := app.Listen(":"+internal.Config.Port, fiber.ListenConfig{
		GracefulContext: ctx,
		ShutdownTimeout: 10 * time.Second,
	})
	if err != nil {
		slog.Error("server stopped", "error", err)
	}
}
