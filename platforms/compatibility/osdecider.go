package compatibility

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"github.com/EdgeLordKirito/wallpapersetter/internal/appinfo"
)

func GetCurrentOS() string {
	return runtime.GOOS
}

func GetAppConfigPath() string {
	return filepath.Join(GetAppConfigDir(), appinfo.AppName+".toml")

}

func GetAppConfigDir() string {
	configDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatalf("Error determining config directory: %v\n", err)
	}
	return filepath.Join(configDir, appinfo.AppName)
}

func GetAppExtensionConfigDir() string {
	return filepath.Join(GetAppConfigDir(), "extensions")
}

func GetAppCacheDir() string {
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		log.Fatalf("Error determining cache directory: %v\n", err)
	}
	return filepath.Join(cacheDir, appinfo.AppName)
}

func GetURLCacheDir() string {
	dir := GetAppCacheDir()
	return filepath.Join(dir, "url")
}

func GetAppBitmapCacheDir() string {
	return path.Join(GetAppCacheDir(), "bitmaps")
}
