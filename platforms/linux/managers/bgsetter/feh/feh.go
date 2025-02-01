package feh

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
	conf := loadFehConfig()
	abs, err := filepath.Abs(input)
	if err != nil {
		log.Fatalf("Failed to get absolute path: %v", err)
	}
	commandString := buildFehCommand(conf, abs)
	cmd := exec.Command("bash", "-c", commandString)
	err = cmd.Run()

	if err != nil {
		log.Printf("Error setting wallpaper: %v", err)
		return fmt.Errorf("failed to set wallpaper: %w", err)
	}

	log.Printf("Wallpaper set to: file://%s", input)
	return nil
}

func buildFehCommand(conf fehConfig, imagePath string) string {
	if conf.Mode == "" {
		conf.Mode = "scale" // Default mode
	}

	// Escape any necessary characters in the path
	escapedPath := strings.ReplaceAll(imagePath, "'", "'\\''")

	// Construct the command
	return fmt.Sprintf("feh --bg-%s '%s'", conf.Mode, escapedPath)
}

type fehConfig struct {
	Mode string `toml:"scaling_mode"`
}

func loadFehConfig() fehConfig {
	configDir := compatibility.GetAppExtensionConfigDir()
	entries, err := os.ReadDir(configDir)
	if err != nil {
		log.Printf("Failed to read extensions config directory: %v", err)
		return fehConfig{} // Return default config if directory can't be read
	}

	var configPath string
	for _, entry := range entries {
		if !entry.IsDir() && strings.EqualFold(entry.Name(), "feh.toml") {
			configPath = path.Join(configDir, entry.Name())
			break
		}
	}

	if configPath == "" {
		log.Println("No feh.toml configuration file found, using defaults")
		return fehConfig{}
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		log.Printf("Failed to read feh.toml: %v", err)
		return fehConfig{} // Return default if file can't be read
	}

	var config fehConfig
	if err := toml.Unmarshal(data, &config); err != nil {
		log.Printf("Failed to parse feh.toml: %v", err)
		return fehConfig{} // Return default if parsing fails
	}

	return config
}
