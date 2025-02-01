package windows

import "github.com/EdgeLordKirito/wallpapersetter/internal/config"

func GetBackendStrategy(conf config.Config) func(string) error {
	return setWallpaperWinApi
}
