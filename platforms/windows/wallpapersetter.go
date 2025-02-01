//go:build windows

package windows

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"log"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"unsafe"

	_ "image/gif"  // Import GIF support
	_ "image/jpeg" // Import JPEG support
	_ "image/png"  // Import PNG support

	"github.com/EdgeLordKirito/wallpapersetter/internal/imagevalidator"
	"github.com/EdgeLordKirito/wallpapersetter/platforms/compatibility"
	"golang.org/x/image/bmp"
)

// Constants for SystemParametersInfoW
const (
	win_SPI_SETDESKWALLPAPER = 0x0014
	win_SPIF_UPDATEINIFILE   = 0x01
	win_SPIF_SENDCHANGE      = 0x02
)

// SetWallpaper sets the desktop wallpaper in Windows
func SetWallpaper(input string) error {
	wallpaperPath, err := wallpaperPrechecks(input)
	// Windows API call to set wallpaper
	user32 := syscall.NewLazyDLL("user32.dll")
	systemParametersInfo := user32.NewProc("SystemParametersInfoW")

	lpwsz, err := syscall.UTF16PtrFromString(wallpaperPath)
	if err != nil {
		return fmt.Errorf("failed to convert path to UTF-16: %w", err)
	}

	ret, _, err := systemParametersInfo.Call(
		win_SPI_SETDESKWALLPAPER,
		0,
		uintptr(unsafe.Pointer(lpwsz)),
		win_SPIF_UPDATEINIFILE|win_SPIF_SENDCHANGE,
	)

	if ret == 0 {
		return fmt.Errorf("failed to set wallpaper: %v", err)
	}

	return nil
}

func wallpaperPrechecks(input string) (string, error) {
	if !imagevalidator.IsImageFile(input) {
		return "", imagevalidator.ErrNotAnImage
	}
	ext := strings.ToLower(filepath.Ext(input))

	if ext != ".bmp" {
		log.Printf("Input is not a BMP: '%s'", input)

		// Check the cache first to see if BMP already exists
		cachedPath, err := checkCacheForBMP(input)
		if err == nil {
			return cachedPath, nil
		}

		// If not found in cache, convert the image to BMP
		convertedPath, err := convertToBMP(input)
		if err != nil {
			log.Printf("Failed to convert image to BMP: '%s'", convertedPath)
			return "", fmt.Errorf("failed to convert image to BMP: %w", err)
		}
		// Use the cached or newly converted BMP path
		log.Printf("Converted input File to BMP and saved it at: '%s'", convertedPath)
		return convertedPath, nil
	}
	return input, nil
}

func convertToBMP(inputPath string) (string, error) {
	// Open the input image file
	file, err := os.Open(inputPath)
	if err != nil {

		log.Printf("failed to open image: %v", err)
		return "", fmt.Errorf("failed to open image: %w", err)
	}
	defer file.Close()

	// Decode image (supports JPEG, PNG, GIF, TIFF, WebP, HEIF, etc.)
	img, format, err := image.Decode(file)
	if err != nil {
		log.Printf("failed to decode image: %v", err)
		return "", fmt.Errorf("failed to decode image: %w", err)
	}

	// Convert to BMP-compatible format (RGBA)
	rgba := image.NewRGBA(img.Bounds())
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{}, draw.Src)

	// Handle transparency in PNG (ensure fully opaque background)
	if format == "png" {
		for y := 0; y < rgba.Bounds().Dy(); y++ {
			for x := 0; x < rgba.Bounds().Dx(); x++ {
				c := rgba.At(x, y)
				r, g, b, _ := c.RGBA()
				rgba.Set(x, y, color.RGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), 255})
			}
		}
	}

	// Get the config directory path from GetConfigPath()
	configDir := compatibility.GetAppBitmapCacheDir()
	if configDir == "" {
		return "", fmt.Errorf("could not get config directory path")
	}

	// Generate BMP file path in the config directory
	filenameWithoutExt := strings.TrimSuffix(filepath.Base(inputPath), filepath.Ext(inputPath))
	bmpPath := filepath.Join(configDir, filenameWithoutExt+".bmp")

	// Ensure the config directory exists
	if err := os.MkdirAll(configDir, os.ModePerm); err != nil {
		return "", fmt.Errorf("failed to create config directory: %w", err)
	}

	// Save image as BMP, overwriting if necessary
	bmpFile, err := os.Create(bmpPath)
	if err != nil {
		log.Printf("failed to create BMP file: %v", err)
		return "", fmt.Errorf("failed to create BMP file: %w", err)
	}
	defer bmpFile.Close()

	if err := bmp.Encode(bmpFile, rgba); err != nil {
		log.Printf("failed to encode BMP: %v", err)
		return "", fmt.Errorf("failed to encode BMP: %w", err)
	}

	return bmpPath, nil
}

func checkCacheForBMP(inputPath string) (string, error) {
	// Get the config directory where BMPs are cached
	cacheDir := compatibility.GetAppBitmapCacheDir()

	// Generate BMP file path in the cache directory
	filenameWithoutExt := strings.TrimSuffix(filepath.Base(inputPath), filepath.Ext(inputPath))
	bmpPath := filepath.Join(cacheDir, filenameWithoutExt+".bmp")

	// Check if the BMP file already exists in the cache
	if _, err := os.Stat(bmpPath); err == nil {
		// BMP file exists in cache, return the cached path
		log.Printf("Using cached BMP file: '%s'", bmpPath)
		return bmpPath, nil
	}

	// BMP not found in cache, return error to indicate conversion is required
	return "", fmt.Errorf("BMP not found in cache, conversion needed")
}
