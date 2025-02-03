package windows

import "github.com/EdgeLordKirito/wallpapersetter/internal/config"

func GetBackendStrategy(conf *config.Windows) func(string) error {
	return setWallpaperWinApi
}
