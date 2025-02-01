package gnome

import (
	"fmt"
	"log"
	"os/exec"
	"path/filepath"

	"github.com/EdgeLordKirito/wallpapersetter/internal/filevalidator"
)

func SetWallpaper(input string) error {
	abs, err := filepath.Abs(input)
	if err != nil {
		log.Fatalf("Failed to get absolute path: %v", err)
	}
	if err := filevalidator.IsValidFile(abs); err != nil {
		return fmt.Errorf("Unable to set Wallpaper '%s' reason '%v'", abs, err)
	}

	uri := "file://" + abs

	// Set the wallpaper using gsettings
	args := []string{
		"set", "org.gnome.desktop.background",
		"picture-uri", filevalidator.QuotePath(uri),
	}
	cmd := exec.Command("gsettings", args...)
	err = cmd.Run()
	if err != nil {
		log.Printf("Error setting wallpaper: %v", err) // Log the error
		return fmt.Errorf("failed to set wallpaper: %w", err)
	}

	log.Printf("Wallpaper set to: %s", uri)
	return nil
}
