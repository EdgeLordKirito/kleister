package independent

import (
	"errors"
	"fmt"
	"strings"

	"github.com/EdgeLordKirito/wallpapersetter/internal/config"
	"github.com/EdgeLordKirito/wallpapersetter/platforms/compatibility"
	"github.com/EdgeLordKirito/wallpapersetter/platforms/darwin"
	"github.com/EdgeLordKirito/wallpapersetter/platforms/linux"
	"github.com/EdgeLordKirito/wallpapersetter/platforms/windows"
)

var (
	ErrUnsupportedOS error = errors.New("Encountered Unsupported OS")
)

func GetBackendStrategy(conf *config.MainConfig) WallpaperSetter {
	os := strings.ToLower(compatibility.GetCurrentOS())
	var strategy SetWallpaperFunc
	switch os {
	case "windows":
		strategy = windows.GetBackendStrategy(&conf.Windows)
	case "darwin":
		strategy = darwin.GetBackendStrategy(&conf.Darwin)
	case "linux":
		strategy = linux.GetBackendStrategy(&conf.Linux)
	default:
		strategy = unsupported(os)
	}
	return strategy
}

func unsupported(os string) func(string) error {
	return func(string) error {
		return fmt.Errorf(
			"Encountered Unsupported Operating System '%s'",
			os)
	}
}
