package bgsetter

import (
	"fmt"
	"strings"

	"github.com/EdgeLordKirito/wallpapersetter/internal/config"
	"github.com/EdgeLordKirito/wallpapersetter/platforms/linux/strategies/bgsetter/feh"
	"github.com/EdgeLordKirito/wallpapersetter/platforms/linux/strategies/bgsetter/nitrogen"
)

func GetBackendStrategy(conf config.Config) func(string) error {
	lower := strings.ToLower(conf.Backend)
	switch lower {
	case "feh":
		return feh.SetWallpaper
	case "nitrogen":
		return nitrogen.SetWallpaper
	default:
		return unsupported(lower)
	}
}

func unsupported(de string) func(string) error {
	return func(string) error {
		return fmt.Errorf(
			"Encountered Unsupported Background Setter '%s'",
			de)
	}
}
