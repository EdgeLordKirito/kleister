package feh

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/EdgeLordKirito/wallpapersetter/internal/filevalidator"
	"github.com/EdgeLordKirito/wallpapersetter/platforms/compatibility"
	"github.com/pelletier/go-toml/v2"
)

func SetWallpaper(input string) error {
	abs, err := filepath.Abs(input)
	if err != nil {
		log.Fatalf("Failed to get absolute path: %v", err)
	}
	if err := filevalidator.IsValidFile(abs); err != nil {
		return fmt.Errorf("Unable to set Wallpaper '%s' reason '%v'", abs, err)
	}
	conf := loadConfig()
	args := []string{
		fmt.Sprintf("--bg-%s", conf.getMode()), filevalidator.QuotePath(abs),
	}
	cmd := exec.Command("feh", args...)
	err = cmd.Run()

	if err != nil {
		log.Printf("Error setting wallpaper: %v", err)
		return fmt.Errorf("failed to set wallpaper: %w", err)
	}

	log.Printf("Wallpaper set to: file://%s", input)
	return nil
}

type config struct {
	Mode string `toml:"scaling_mode"`
}

func (self config) getMode() string {
	if self.Mode == "" {
		return "scale"
	}
	return self.Mode
}

func loadConfig() config {
	configDir := compatibility.GetAppExtensionConfigDir()
	entries, err := os.ReadDir(configDir)
	if err != nil {
		log.Printf("Failed to read extensions config directory: %v", err)
		return config{} // Return default config if directory can't be read
	}

	var configPath string
	for _, entry := range entries {
		if !entry.IsDir() && strings.EqualFold(entry.Name(), "feh.toml") {
			configPath = filepath.Join(configDir, entry.Name())
			break
		}
	}

	if configPath == "" {
		log.Println("No feh.toml configuration file found, using defaults")
		return config{}
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		log.Printf("Failed to read feh.toml: %v", err)
		return config{} // Return default if file can't be read
	}

	var config config
	if err := toml.Unmarshal(data, &config); err != nil {
		log.Printf("Failed to parse feh.toml: %v", err)
		return config // Return default if parsing fails
	}

	return config
}
