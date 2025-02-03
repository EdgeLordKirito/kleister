package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/EdgeLordKirito/wallpapersetter/platforms/compatibility"
	"github.com/pelletier/go-toml/v2"
)

type Config interface {
	Dirs() []string
	Backend() string
}

type MainConfig struct {
	Windows Windows `toml:"windows"`
	Linux   Linux   `toml:"linux"`
	Darwin  Darwin  `toml:"darwin"`
}

func (self MainConfig) GetOSConfig() Config {
	os := strings.ToLower(compatibility.GetCurrentOS())
	switch os {
	case "windows":
		return self.Windows
	case "darwin":
		return self.Darwin
	case "linux":
		return self.Linux
	default:
		return nil
	}
}

type Windows struct {
	WallpaperDirs []string `toml:"wallpapers"`
	ChosenBackend string   `toml:"backend"`
}

func (self Windows) Dirs() []string {
	return self.WallpaperDirs
}

func (self Windows) Backend() string {
	return self.ChosenBackend
}

type Linux struct {
	WallpaperDirs []string `toml:"wallpapers"`
	ChosenBackend string   `toml:"backend"`
}

func (self Linux) Dirs() []string {
	return self.WallpaperDirs
}

func (self Linux) Backend() string {
	return self.ChosenBackend
}

type Darwin struct {
	WallpaperDirs []string `toml:"wallpapers"`
	ChosenBackend string   `toml:"backend"`
}

func (self Darwin) Dirs() []string {
	return self.WallpaperDirs
}

func (self Darwin) Backend() string {
	return self.ChosenBackend
}

func Marshal(config MainConfig) ([]byte, error) {
	data, err := toml.Marshal(config)
	if err != nil {
		return nil, fmt.Errorf("error marshalling TOML: %w", err)
	}
	return data, nil
}

func Unmarshal(data []byte) (MainConfig, error) {
	var config MainConfig
	err := toml.Unmarshal(data, &config)
	if err != nil {
		return config, fmt.Errorf("error unmarshalling TOML: %w", err)
	}
	return config, nil
}

func GetUserConfig() (*MainConfig, error) {
	path := compatibility.GetAppConfigPath()

	// Check if the file exists
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("config file does not exist: %s", path)
		}
		return nil, fmt.Errorf("error checking config file: %w", err)
	}

	// Read file contents
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	// Unmarshal the TOML data into Config struct
	config, err := Unmarshal(data)
	if err != nil {
		return nil, fmt.Errorf("error parsing config file: %w", err)
	}

	return &config, nil
}
