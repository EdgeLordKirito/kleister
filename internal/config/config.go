package config

import (
	"fmt"
	"os"

	"github.com/EdgeLordKirito/wallpapersetter/platforms/compatibility"
	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	//TODO: make this an array for multiple dir support
	WallpaperDirs []string `toml:"wallpapers"`
	Backend       string   `toml:"backend"`
}

func Marshal(config Config) ([]byte, error) {
	data, err := toml.Marshal(config)
	if err != nil {
		return nil, fmt.Errorf("error marshalling TOML: %w", err)
	}
	return data, nil
}

func Unmarshal(data []byte) (Config, error) {
	var config Config
	err := toml.Unmarshal(data, &config)
	if err != nil {
		return config, fmt.Errorf("error unmarshalling TOML: %w", err)
	}
	return config, nil
}

func GetUserConfig() (*Config, error) {
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
