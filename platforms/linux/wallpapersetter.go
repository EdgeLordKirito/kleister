package linux

import (
	"errors"

	"github.com/EdgeLordKirito/wallpapersetter/platforms/linux/managers/bgsetter"
	"github.com/EdgeLordKirito/wallpapersetter/platforms/linux/managers/desktop"
)

var (
	ErrUnsupportedDEIdentifier error = errors.New("Unsupported Desktop Environment Identifier")
)

func SetWallpaper(input string) error {
	env := getDesktopEnvironment()
	envType := getEnvironmentType(env)

	switch envType {
	case DesktopEnvironment:
		return desktop.Redirect(env, input)
	case WindowManager:
		return bgsetter.Redirect(env, input)
	default:
		return ErrUnsupportedDEIdentifier
	}
}
