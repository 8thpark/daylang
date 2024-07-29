package main

import (
	_ "embed"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pkg/errors"
)

//go:embed VERSION
var version string

func main() {
	router := echo.New()

	router.HideBanner = true
	router.Use(defaultMiddleware()...)

	router.GET("/logout", func(ctx echo.Context) error {
		log.Println("user logged out")
		return ctx.Redirect(http.StatusSeeOther, "/?logout=true")
	})

	router.GET("/status", func(ctx echo.Context) error {
		status := map[string]string{
			"status":  "ok",
			"version": strings.TrimSpace(version),
		}

		return ctx.JSON(http.StatusOK, status)
	})

	router.GET("/*", func(ctx echo.Context) error {
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
	})

	addr := address()
	router.Logger.Fatal(router.Start(addr))
}

func address() string {
	host := "localhost"
	if os.Getenv("APP_ENV") == "docker" {
		host = ""
	}

	return fmt.Sprintf("%s:5174", host)
}

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
