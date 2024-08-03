package config

import (
	"fmt"
	"strings"
)

// Environment is an enum for the different environments the application can run in.
type Environment string

const (
	EnvironmentDocker Mode = "docker"
	EnvironmentLocal  Mode = "local"
)

func environments() []Mode {
	return []Mode{
		EnvironmentDocker,
		EnvironmentLocal,
	}
}

func joinEnvironments() string {
	environmentsList := environments()
	environmentStr := make([]string, 0, len(environmentsList))

	for _, e := range environmentsList {
		environmentStr = append(environmentStr, fmt.Sprintf(`"%s"`, e))
	}

	return strings.Join(environmentStr, ", ")
}
