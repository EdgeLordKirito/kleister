package linux

import "os"

func init() {
	desktopEnvironmentSet = make(map[string]bool)
	for _, env := range desktopEnvironments {
		desktopEnvironmentSet[env] = true // Add each desktop environment to the set
	}

	windowManagerSet = make(map[string]bool)
	for _, wm := range windowManagers {
		windowManagerSet[wm] = true // Add each window manager to the set
	}
}

type EnvironmentType int

const (
	DesktopEnvironment EnvironmentType = iota
	WindowManager
	Unknown
)

func getEnvironmentType(env string) EnvironmentType {
	if desktopEnvironmentSet[env] {
		return DesktopEnvironment
	} else if windowManagerSet[env] {
		return WindowManager
	} else {
		return Unknown
	}
}

var desktopEnvironments = []string{
	"GNOME", "KDE", "XFCE", "Mate", "Cinnamon", "Unity",
	"LXQt",
}

var desktopEnvironmentSet map[string]bool

func DesktopEnvironments() []string {
	return append([]string(nil), desktopEnvironments...)
}

func DesktopEnvironmentSet() map[string]bool {
	// Creating a new map to return a copy of the original set
	copySet := make(map[string]bool)
	for k, v := range desktopEnvironmentSet {
		copySet[k] = v
	}
	return copySet
}

var windowManagers = []string{
	"i3", "Openbox", "Awesome", "Fluxbox", "Xmonad", "LeftWM",
	"bspwm", "Herbstluftwm", "Ratpoison",
}

var windowManagerSet map[string]bool

func WindowManagers() []string {
	return append([]string(nil), windowManagers...)
}

func WindowManagersSet() map[string]bool {
	// Creating a new map to return a copy of the original set
	copySet := make(map[string]bool)
	for k, v := range windowManagerSet {
		copySet[k] = v
	}
	return copySet
}

// Helper function to check if a value exists in a set
func containsSet(set map[string]bool, value string) bool {
	return set[value]
}

func getDesktopEnvironment() string {
	// First, try to get the DE environment variable
	de := os.Getenv("DE")
	if de != "" {
		return de
	}

	// If DE is not set, try XDG_CURRENT_DESKTOP
	de = os.Getenv("XDG_CURRENT_DESKTOP")
	if de != "" {
		return de
	}

	// If XDG_CURRENT_DESKTOP is also empty, try XDG_SESSION_DESKTOP
	de = os.Getenv("XDG_SESSION_DESKTOP")
	if de != "" {
		return de
	}

	// If none of them are set, return an empty string or a default message
	return "Unknown"
}
