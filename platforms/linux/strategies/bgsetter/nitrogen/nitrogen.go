package nitrogen

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
	args := []string{"--set", filevalidator.QuotePath(abs)}
	cmd := exec.Command("nitrogen", args...)
	err = cmd.Run()

	if err != nil {
		log.Printf("Error setting wallpaper: %v", err)
		return fmt.Errorf("failed to set wallpaper: %w", err)
	}

	log.Printf("Wallpaper set to: file://%s", input)
	return nil

}
