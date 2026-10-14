package internal

import (
	"errors"
	"log/slog"
	"time"

	"github.com/gofiber/fiber/v3"

	"booking-api/utils"
)

// NewApp builds the Fiber app and registers every route.
func NewApp() *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:      "booking-api",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
		ErrorHandler: handleError,
	})

	app.Get("/health", handleHealth)

	return app
}

// handleError turns any error returned by a handler into our standard
// JSON shape: { "error": { "code": ..., "message": "..." } }
func handleError(c fiber.Ctx, err error) error {
	status := fiber.StatusInternalServerError
	message := "something went wrong"

	var fe *fiber.Error
	if errors.As(err, &fe) {
		status = fe.Code
		message = fe.Message
	} else {
		// Unexpected error: log the details, but don't leak them to the client.
		slog.Error("unhandled error", "method", c.Method(), "path", c.Path(), "error", err)
	}

	return c.Status(status).JSON(utils.ErrorResponse{
		Error: utils.ErrorBody{Code: status, Message: message},
	})
}
