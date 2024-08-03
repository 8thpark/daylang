package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func defaultMiddleware() []echo.MiddlewareFunc {
	return []echo.MiddlewareFunc{
		middleware.Gzip(),
		middleware.Logger(),
		middleware.Recover(),
		noIndexMiddleware,
	}
}

func noIndexMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("X-Robots-Tag", "noindex")
		return next(c)
	}
}
