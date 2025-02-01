package gnome

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/EdgeLordKirito/wallpapersetter/internal/imagevalidator"
)

func SetWallpaper(input string) error {
	if !imagevalidator.IsImageFile(input) {
		return imagevalidator.ErrNotAnImage
	}
	// Check if the input file exists
	if _, err := os.Stat(input); err != nil {
		return fmt.Errorf("error getting file: %s", input)
	}

	// Convert the input file path to a file URI
	absPath, err := filepath.Abs(input)
	if err != nil {
		return fmt.Errorf("failed to get absolute path: %w", err)
	}

	uri := "file://" + absPath

	// Set the wallpaper using gsettings
	cmd := exec.Command("gsettings", "set", "org.gnome.desktop.background", "picture-uri", uri)
	err = cmd.Run()
	if err != nil {
		log.Printf("Error setting wallpaper: %v", err) // Log the error
		return fmt.Errorf("failed to set wallpaper: %w", err)
	}

	log.Printf("Wallpaper set to: %s", uri)
	return nil
}
