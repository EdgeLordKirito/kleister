package compatibility

import (
	"log"
	"os"
	"path"
	"runtime"

	"github.com/EdgeLordKirito/wallpapersetter/internal/appinfo"
)

func GetCurrentOs() string {
	return runtime.GOOS
}

func GetAppConfigPath() string {
	return path.Join(GetAppConfigDir(), appinfo.AppName+".toml")

}

func GetAppConfigDir() string {
	configDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatalf("Error determining config directory: %v\n", err)
	}
	return path.Join(configDir, appinfo.AppName)
}

func GetAppExtensionConfigDir() string {
	return path.Join(GetAppConfigDir(), "extensions")
}

func GetAppCacheDir() string {
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		log.Fatalf("Error determining cache directory: %v\n", err)
	}
	return path.Join(cacheDir, appinfo.AppName)
}

func GetAppBitmapCacheDir() string {
	return path.Join(GetAppCacheDir(), "bitmaps")
}
