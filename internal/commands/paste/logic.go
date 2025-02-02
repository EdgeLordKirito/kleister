package paste

import (
	"errors"
	"fmt"
	"os"

	"github.com/EdgeLordKirito/wallpapersetter/internal/config"
	"github.com/EdgeLordKirito/wallpapersetter/internal/filevalidator"
	"github.com/EdgeLordKirito/wallpapersetter/platforms/independent"
	"github.com/spf13/cobra"
)

var (
	ErrInvalidPathType error = errors.New(
		"Invalid Pathtype expecting valid File Path or Directory Path or valid Url")
	ErrUnimplementedUrl error = errors.New("Url support is not yet implemented")
)

func Run(cmd *cobra.Command, args []string) error {
	conf, err := config.GetUserConfig()
	if err != nil {
		return fmt.Errorf("Unable to read Config %v", err)
	}
	strategy := independent.GetBackendStrategy(conf)
	pathType := determinePathType(path)
	switch pathType {
	case "directory":
		return handleDirectory(conf, strategy)
	case "file":
		return handleFile(conf, strategy)
	case "url":
		return handleUrl(conf, strategy)
	default:
		return ErrInvalidPathType
	}
}

func handleDirectory(conf *config.Config, strategy independent.WallpaperSetter) error {
	//either load the dir or load the config dirs
	var dirs []string = conf.WallpaperDirs
	if path != "" {
		dirs = []string{path}
	}
	_ = dirs
	//build the iterator
	file := ""
	return strategy.Set(file)
}

func handleFile(conf *config.Config, strategy independent.WallpaperSetter) error {
	_ = conf
	return strategy.Set(path)
}

func handleUrl(conf *config.Config, strategy independent.WallpaperSetter) error {
	// TODO: download the file from the url
	_, _ = conf, strategy
	return ErrUnimplementedUrl
}

func determinePathType(input string) string {
	// If input is empty, assume it's a directory
	if input == "" {
		return "directory"
	}

	// Check if input is a valid URL
	if filevalidator.IsValidURL(input) {
		return "url"
	}

	// Check if it's an existing file or directory
	info, err := os.Stat(input)
	if err == nil {
		if info.IsDir() {
			return "directory"
		} else {
			return "file"
		}
	}

	// If nothing matched, it's invalid
	return "invalid"
}
