package main

import (
	"fmt"

	"github.com/EdgeLordKirito/wallpapersetter/internal/config"
	"github.com/EdgeLordKirito/wallpapersetter/platforms/independent"
)

func main() {
	//wallpaper := `C:\Users\Till\Pictures\Wallpapers\Wallpapers\MikuNatureFramed.jpg`

	wallpaper := `C:\Users\Till\Pictures\Wallpapers\Wallpapers\Fern Nest.jpg`
	conf := config.Config{
		WallpaperDir: "",
		Backend:      "windows"}
	var ws independent.WallpaperSetter
	ws = independent.GetBackendStrategy(conf)
	err := ws.Set(wallpaper)
	fmt.Println(err)
}
