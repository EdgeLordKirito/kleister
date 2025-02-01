package darwin

import (
	"fmt"
	"log"
	"os/exec"
)

func SetWallpaper(input string) error {
	// Prepare the AppleScript to set the wallpaper
	appleScript := fmt.Sprintf(`tell application "System Events"
		set desktop picture to POSIX file "%s"
	end tell`, input)

	// Execute the osascript command
	cmd := exec.Command("osascript", "-e", appleScript)
	err := cmd.Run()

	// If there was an error, log it and return
	if err != nil {
		// Log the error before returning
		log.Printf("Error occurred while setting wallpaper: %v", err)
		return fmt.Errorf("failed to set wallpaper on macOS: %w", err)
	}

	log.Printf("Wallpaper set to: %s", input)
	return nil
}
