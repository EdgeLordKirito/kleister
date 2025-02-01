package darwin

import (
	"fmt"
	"log"
	"os/exec"
	"path/filepath"

	"github.com/EdgeLordKirito/wallpapersetter/internal/filevalidator"
)

func setWallpaperOSA(input string) error {
	abs, err := filepath.Abs(input)
	if err != nil {
		log.Fatalf("Failed to get absolute path: %v", err)
	}
	if err := filevalidator.IsValidFile(abs); err != nil {
		return fmt.Errorf("Unable to set Wallpaper '%s' reason '%v'", abs, err)
	}
	// Prepare the AppleScript to set the wallpaper
	appleScript := fmt.Sprintf(`tell application "System Events"
		set desktop picture to POSIX file "%s"
	end tell`, input)
	// Execute the osascript command
	cmd := exec.Command("osascript", "-e", appleScript)
	err = cmd.Run()

	// If there was an error, log it and return
	if err != nil {
		// Log the error before returning
		log.Printf("Error occurred while setting wallpaper: %v", err)
		return fmt.Errorf("failed to set wallpaper on macOS: %w", err)
	}

	log.Printf("Wallpaper set to: %s", input)
	return nil
}
