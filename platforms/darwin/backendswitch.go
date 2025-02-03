package darwin

import "github.com/EdgeLordKirito/wallpapersetter/internal/config"

func GetBackendStrategy(conf *config.Darwin) func(string) error {
	return setWallpaperOSA
}
