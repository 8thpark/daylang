package router

import (
	"fmt"

	"github.com/8thpark/daylang/server/config"
	"github.com/8thpark/daylang/server/handler"
	"github.com/labstack/echo/v4"
)

// New creates a new router.
func New(cfg *config.Config) *echo.Echo {
	rtr := echo.New()

	rtr.HideBanner = true
	rtr.Use(defaultMiddleware()...)

	// API routes.
	rtr.GET(handler.PathSelf, handler.Self(cfg))

	// Utility routes.
	rtr.GET(handler.PathStatus, handler.Status(cfg))

	// Route to serve the React app, must be last.
	rtr.GET(handler.PathWildcard, handler.Wildcard(cfg))

	return rtr
}

func Start(cfg *config.Config, rtr *echo.Echo) {
	addr := address(cfg)
	rtr.Logger.Fatal(rtr.Start(addr))
}

func address(cfg *config.Config) string {
	host := "localhost"
	if cfg.Environment == string(config.EnvironmentDocker) {
		host = ""
	}

	return fmt.Sprintf("%s:5174", host)
}
