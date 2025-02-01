package config

import (
	"fmt"

	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	WallpaperDir string `toml:"wallpapers"`
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
