package linux

import (
	"github.com/EdgeLordKirito/wallpapersetter/internal/config"
	"github.com/EdgeLordKirito/wallpapersetter/platforms/linux/strategies/bgsetter"
	"github.com/EdgeLordKirito/wallpapersetter/platforms/linux/strategies/desktop"
)

func GetBackendStrategy(conf config.Config) func(string) error {
	if conf.Backend != "" {
		return bgsetter.GetBackendStrategy(conf)
	}
	return desktop.GetDEBackendStrategy(conf)
}
