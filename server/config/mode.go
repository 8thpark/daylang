package config

import (
	"fmt"
	"strings"
)

// Mode is an enum for the different modes the application can run in.
type Mode string

const (
	ModeBridge     Mode = "bridge"
	ModeStandalone Mode = "standalone"
)

func modes() []Mode {
	return []Mode{
		ModeBridge,
		ModeStandalone,
	}
}

func joinModes() string {
	modesList := modes()
	modeStr := make([]string, 0, len(modesList))

	for _, m := range modesList {
		modeStr = append(modeStr, fmt.Sprintf(`"%s"`, m))
	}

	return strings.Join(modeStr, ", ")
}
