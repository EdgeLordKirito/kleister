//go:build windows

package windows

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/EdgeLordKirito/wallpapersetter/internal/filevalidator"
)

// Constants for SystemParametersInfoW
const (
	win_SPI_SETDESKWALLPAPER = 0x0014
	win_SPIF_UPDATEINIFILE   = 0x01
	win_SPIF_SENDCHANGE      = 0x02
)

// setWallpaperWinApi sets the desktop wallpaper in Windows
func setWallpaperWinApi(input string) error {
	err := filevalidator.IsValidFile(input)
	if err != nil {
		return fmt.Errorf("Unable to set Wallpaper '%s' reason '%v'", input, err)
	}
	// Windows API call to set wallpaper
	user32 := syscall.NewLazyDLL("user32.dll")
	systemParametersInfo := user32.NewProc("SystemParametersInfoW")

	lpwsz, err := syscall.UTF16PtrFromString(input)
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
