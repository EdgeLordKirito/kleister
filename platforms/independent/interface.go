package independent

import ()

// WallpaperSetter defines the contract that any type (including function types) must implement
// to set a wallpaper.
type WallpaperSetter interface {
	Set(input string) error
}

// Define a function type that matches the signature of the Set method
type SetWallpaperFunc func(input string) error

// Implement the WallpaperSetter interface for SetWallpaperFunc
func (f SetWallpaperFunc) Set(input string) error {
	return f(input)
}
