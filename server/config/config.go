package config

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

// Config holds the configuration for the application.
type Config struct {
	BillingRedirect string
	Environment     string
	Mode            string
	Version         string
}

// New initialises a new Config struct.
func New(version string) (*Config, error) {
	return &Config{
		Version: strings.TrimSpace(version),
	}, nil
}

// Flags returns CLI flags for the application, that map to the Config struct.
func (c *Config) Flags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "mode",
			Usage:       fmt.Sprintf("Mode to run the application in (%s).", joinModes()),
			EnvVars:     []string{"DAYLANG_MODE"},
			Value:       string(ModeStandalone),
			Destination: &c.Mode,
			Hidden:      true,
			Action: func(ctx *cli.Context, mode string) error {
				for _, m := range modes() {
					if string(m) == mode {
						return nil
					}
				}

				return errors.Errorf("invalid mode, valid modes: %s", joinModes())
			},
		},
		&cli.StringFlag{
			Name:        "environment",
			Usage:       "Defines runtime behaviors based on the environment.",
			EnvVars:     []string{"DAYLANG_ENVIRONMENT"},
			Destination: &c.Environment,
			Required:    true,
			Hidden:      true,
			Action: func(ctx *cli.Context, environment string) error {
				for _, e := range environments() {
					if string(e) == environment {
						return nil
					}
				}

				return errors.Errorf("invalid environment, valid environments: %s", joinEnvironments())
			},
		},
		&cli.StringFlag{
			Name:        "billing-redirect",
			Usage:       "URL to redirect to when billing is required.",
			EnvVars:     []string{"DAYLANG_BILLING_REDIRECT"},
			Destination: &c.BillingRedirect,
			Hidden:      true,
			Action: func(ctx *cli.Context, uri string) error {
				if uri == "" {
					return nil
				}

				if _, err := url.ParseRequestURI(uri); err != nil {
					return errors.Wrap(err, "billing redirect is not a valid url")
				}

				return nil
			},
		},
	}
}

// Validate checks that all flag values are valid. Most of the time this can be done within the
// Action() function of the flag, but sometimes it's necessary to validate based on the state of
// other flags.
func (c Config) Validate() error {
	if c.Mode == string(ModeBridge) && c.BillingRedirect == "" {
		return errors.New("billing redirect url is required when running in bridge mode")
	}

	return nil
}
