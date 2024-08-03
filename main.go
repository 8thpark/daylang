package main

import (
	_ "embed"
	"log"
	"os"

	"github.com/8thpark/daylang/server/config"
	"github.com/8thpark/daylang/server/router"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

//go:embed VERSION
var version string

func main() {
	cfg, err := config.New(version) // Config values are set within the Action() CLI function.
	if err != nil {
		log.Fatalf("creating config: %v", err)
	}

	app := &cli.App{
		Name:  "daylang",
		Usage: "Daily vocabulary builder for language learners.",
		Flags: cfg.Flags(),
		Action: func(c *cli.Context) error {
			if err := cfg.Validate(); err != nil {
				return errors.Wrap(err, "validating cli flags")
			}

			log.Printf("version: %s", cfg.Version)
			log.Printf("environment: %s", cfg.Environment)
			log.Printf("mode: %s", cfg.Mode)

			if cfg.Mode == string(config.ModeBridge) {
				log.Printf("billing redirect: %s", cfg.BillingRedirect)
			}

			rtr := router.New(cfg)
			router.Start(cfg, rtr)

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatalf("running app: %v", err)
	}
}
