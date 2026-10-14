package internal

import (
	"context"
	"log/slog"
	"time"

	"github.com/gofiber/fiber/v3"

	"booking-api/utils"
)

// handleHealth reports whether the server, Postgres and Redis are working.
// Hosting platforms call this to decide if the app is alive.
func handleHealth(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), 2*time.Second)
	defer cancel()

	res := utils.HealthResponse{Status: "ok", Database: "up", Redis: "up"}

	if err := DB.Ping(ctx); err != nil {
		slog.Error("health check: postgres unreachable", "error", err)
		res.Status, res.Database = "unhealthy", "down"
	}
	if err := Redis.Ping(ctx).Err(); err != nil {
		slog.Error("health check: redis unreachable", "error", err)
		res.Status, res.Redis = "unhealthy", "down"
	}

	if res.Status != "ok" {
		return c.Status(fiber.StatusServiceUnavailable).JSON(res)
	}
	return c.JSON(res)
}
