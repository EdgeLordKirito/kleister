package main

import (
	"fmt"

	"github.com/EdgeLordKirito/wallpapersetter/platforms/independent"
	"github.com/EdgeLordKirito/wallpapersetter/platforms/windows"
)

func main() {
	wallpaper := `C:\Users\Till\Pictures\Wallpapers\Wallpapers\MikuNatureFramed.jpg`

	//wallpaper := `C:\Users\Till\Pictures\Wallpapers\Wallpapers\Fern-Nest.jpg`
	var ws independent.WallpaperSetter
	ws = independent.SetWallpaperFunc(windows.SetWallpaper)
	err := ws.Set(wallpaper)
	fmt.Println(err)
}
