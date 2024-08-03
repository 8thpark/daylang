package handler

import (
	"net/http"

	"github.com/8thpark/daylang/server/config"
	"github.com/labstack/echo/v4"
)

// PathStatus is the route for status handler.
const PathStatus = "/status"

// Status returns the status of the application.
func Status(cfg *config.Config) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		status := map[string]string{
			"status":  "ok",
			"version": cfg.Version,
		}

		return ctx.JSON(http.StatusOK, status)
	}
}
