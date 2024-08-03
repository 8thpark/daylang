package handler

import (
	"os"
	"path/filepath"

	"github.com/8thpark/daylang/server/config"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

// PathWildcard is the route for wildcard handler.
const PathWildcard = "/*"

// Wildcard serves the React app.
func Wildcard(cfg *config.Config) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		path := ctx.Request().URL.Path
		if path == "/" {
			return ctx.File("dist/index.html")
		}

		filePath := filepath.Join("dist", path)

		_, err := os.Stat(filePath)
		if os.IsNotExist(err) { // File doesn't exist, serve index.html.
			return ctx.File("dist/index.html")
		} else if err != nil { // Some other error occurred.
			return errors.Wrap(err, "checking if static file exists")
		}

		return ctx.File(filePath) // File exists, serve it.
	}
}
