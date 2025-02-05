package paste

import (
	"errors"
	"fmt"
	"github.com/EdgeLordKirito/wallpapersetter/internal/config"
	"github.com/EdgeLordKirito/wallpapersetter/internal/filevalidator"
	"github.com/EdgeLordKirito/wallpapersetter/platforms/independent"
	"github.com/spf13/cobra"
	"os"
)

var (
	ErrInvalidPathType error = errors.New(
		"Invalid Pathtype expecting valid File Path or Directory Path or valid Url")
)

func Run(cmd *cobra.Command, args []string) error {
	conf, err := config.GetUserConfig()
	if err != nil {
		return fmt.Errorf("Unable to read Config %v", err)
	}
	pathType := determinePathType(path)
	switch pathType {
	case "directory":
		return handleDirectory(conf)
	case "file":
		return handleFile(conf)
	default:
		return ErrInvalidPathType
	}
}

func handleDirectory(mConf *config.MainConfig) error {
	//either load the dir or load the config dirs
	strategy := independent.GetBackendStrategy(mConf)
	sysConf := mConf.GetOSConfig()
	var dirs []string = sysConf.Dirs()
	if path != "" {
		dirs = []string{path}
	}
	files, _, err := filevalidator.CollectImageFiles(dirs)
	if err != nil {
		return fmt.Errorf(
			"Could not Collect Image files from the directories reason '%v'", err)
	}
	file, err := filevalidator.PickRandomFile(files)
	if err != nil {
		return fmt.Errorf("Could not pick random File reason '%v'", err)
	}
	return strategy.Set(file)
}

func handleFile(mConf *config.MainConfig) error {
	strategy := independent.GetBackendStrategy(mConf)
	return strategy.Set(path)
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
