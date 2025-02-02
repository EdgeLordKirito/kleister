package desktop

import (
	"fmt"
	"os"
	"strings"

	"github.com/EdgeLordKirito/wallpapersetter/internal/config"
	"github.com/EdgeLordKirito/wallpapersetter/platforms/linux/strategies/desktop/gnome"
	"github.com/EdgeLordKirito/wallpapersetter/platforms/linux/strategies/desktop/kde"
)

func unsupported(de string) func(string) error {
	return func(string) error {
		return fmt.Errorf(
			"Encountered Unsupported Desktop Environment '%s'",
			de)
	}
}

func GetDEBackendStrategy(conf *config.Config) func(string) error {
	de := strings.ToLower(getDesktopEnvironment())
	switch de {
	case "gnome":
		return gnome.SetWallpaper
	case "kde":
		return kde.SetWallpaper
	default:
		return unsupported(de)
	}
}

func getDesktopEnvironment() string {
	de := os.Getenv("DE")
	if de != "" {
		return de
	}

	de = os.Getenv("XDG_CURRENT_DESKTOP")
	if de != "" {
		return de
	}

	de = os.Getenv("XDG_SESSION_DESKTOP")
	if de != "" {
		return de
	}

	return "Unknown"
}
