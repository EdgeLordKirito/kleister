package bgsetter

import (
	"strings"

	"github.com/EdgeLordKirito/wallpapersetter/platforms/linux/managers"
	"github.com/EdgeLordKirito/wallpapersetter/platforms/linux/managers/bgsetter/feh"
	"github.com/EdgeLordKirito/wallpapersetter/platforms/linux/managers/bgsetter/nitrogen"
)

func Redirect(identifier, imagePath string) error {
	lower := strings.ToLower(identifier)
	var setterFunction func(input string) error

	switch lower {
	case "feh":
		setterFunction = feh.SetWallpaper
	case "nitrogen":
		setterFunction = nitrogen.SetWallpaper
	default:
		return managers.ErrInvalidIdentifier
	}

	return setterFunction(imagePath)
}
