package imagevalidator

import (
	"errors"
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

var universalImageExtensionsSet = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
}

var ErrNotAnImage error = errors.New("File is not of an Supported Image Type")

// IsImageFile checks if the given image path is a valid image file
func IsImageFile(imagePath string) bool {
	return isExtensionValid(imagePath, imageExtensionsSet)
}

func IsUniversallySupported(imagePath string) bool {
	return isExtensionValid(imagePath, universalImageExtensionsSet)
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
