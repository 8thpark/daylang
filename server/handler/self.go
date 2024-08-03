package handler

import (
	"net/http"

	"github.com/8thpark/daylang/server/config"
	"github.com/labstack/echo/v4"
)

// PathSelf is the route for self handler.
const PathSelf = "/api/self"

// Self returns global configuration state about the application.
func Self(cfg *config.Config) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		status := map[string]string{
			"billingRedirect": cfg.BillingRedirect,
			"mode":            cfg.Mode,
			"version":         cfg.Version,
		}

		return ctx.JSON(http.StatusOK, status)
	}
}
