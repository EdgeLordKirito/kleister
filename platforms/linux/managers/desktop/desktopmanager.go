package desktop

import (
	"strings"

	"github.com/EdgeLordKirito/wallpapersetter/platforms/linux/managers"
	"github.com/EdgeLordKirito/wallpapersetter/platforms/linux/managers/desktop/gnome"
	"github.com/EdgeLordKirito/wallpapersetter/platforms/linux/managers/desktop/kde"
)

func Redirect(identifier, imagePath string) error {
	lower := strings.ToLower(identifier)
	var setterFunction func(input string) error
	switch lower {
	case "gnome":
		setterFunction = gnome.SetWallpaper
	case "kde":
		setterFunction = kde.SetWallpaper
	default:
		return managers.ErrInvalidIdentifier
	}

	return setterFunction(imagePath)
}
