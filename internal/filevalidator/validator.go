package filevalidator

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

var imageExtensionsSet = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
	".bmp":  true,
	".gif":  true,
}

var ErrNotAnImage error = errors.New("File is not of an Supported Image Type")
var ErrPathContainsSpace error = errors.New("Filepath contains atleast one spacecharacter")

// IsImageFile checks if the given image path is a valid image file
func IsImageFile(imagePath string) bool {
	return isExtensionValid(imagePath, imageExtensionsSet)
}

func IsValidFile(input string) error {
	_, err := os.Stat(input)
	if err != nil {
		return err
	}
	if !IsImageFile(input) {
		return ErrNotAnImage
	}

	return nil
}

func isExtensionValid(imagePath string, extSet map[string]bool) bool {
	// Get the file extension (in lowercase to make it case-insensitive)
	ext := strings.ToLower(filepath.Ext(imagePath))

	// Check if the file extension is in the provided set of valid extensions
	if _, valid := extSet[ext]; valid {
		return true
	}

	return false
}

func QuotePath(input string) string {
	// Ensure the input string is enclosed in double quotes.
	return fmt.Sprintf(`"%s"`, input)
}

func IsValidURL(str string) bool {
	parsedURL, err := url.Parse(str)
	return err == nil && parsedURL.Scheme != "" && parsedURL.Host != ""
}

// Checks if a URL ends with a supported image extension
func IsImageURL(str string) bool {
	ext := strings.ToLower(filepath.Ext(str))
	return imageExtensionsSet[ext]
}
