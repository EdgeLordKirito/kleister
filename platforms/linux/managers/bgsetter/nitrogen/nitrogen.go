package nitrogen

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	"github.com/EdgeLordKirito/wallpapersetter/internal/imagevalidator"
	"github.com/EdgeLordKirito/wallpapersetter/platforms/compatibility"
	"github.com/pelletier/go-toml/v2"
)

func SetWallpaper(input string) error {
	if !imagevalidator.IsImageFile(input) {
		return imagevalidator.ErrNotAnImage
	}
	conf := loadNitrogenConfig()
	abs, err := filepath.Abs(input)
	if err != nil {
		log.Fatalf("Failed to get absolute path: %v", err)
	}
	commandString := buildNitrogenCommand(conf, abs)
	cmd := exec.Command("bash", "-c", commandString)
	err = cmd.Run()

	if err != nil {
		log.Printf("Error setting wallpaper: %v", err)
		return fmt.Errorf("failed to set wallpaper: %w", err)
	}

	log.Printf("Wallpaper set to: file://%s", input)
	return nil

}

func buildNitrogenCommand(conf nitrogenConfig, imagePath string) string {
	// Set default mode if not specified
	if conf.Mode == "" {
		conf.Mode = "zoom-fill" // Default mode
	}

	// Escape any necessary characters in the path (for shell compatibility)
	escapedPath := strings.ReplaceAll(imagePath, "'", "'\\''")

	// Construct the nitrogen command
	return fmt.Sprintf("nitrogen --set-%s '%s'", conf.Mode, escapedPath)
}

type nitrogenConfig struct {
	Mode string
}

func loadNitrogenConfig() nitrogenConfig {
	configDir := compatibility.GetAppExtensionConfigDir()
	entries, err := os.ReadDir(configDir)
	if err != nil {
		log.Printf("Failed to read extensions config directory: %v", err)
		return nitrogenConfig{} // Return default config if directory can't be read
	}

	var configPath string
	for _, entry := range entries {
		if !entry.IsDir() && strings.EqualFold(entry.Name(), "nitrogen.toml") {
			configPath = path.Join(configDir, entry.Name())
			break
		}
	}

	if configPath == "" {
		log.Println("No nitrogen.toml configuration file found, using defaults")
		return nitrogenConfig{}
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		log.Printf("Failed to read nitrogen.toml: %v", err)
		return nitrogenConfig{} // Return default if file can't be read
	}

	var config nitrogenConfig
	if err := toml.Unmarshal(data, &config); err != nil {
		log.Printf("Failed to parse nitrogen.toml: %v", err)
		return nitrogenConfig{} // Return default if parsing fails
	}

	return config
}
