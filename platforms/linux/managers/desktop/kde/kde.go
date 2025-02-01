package kde

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/EdgeLordKirito/wallpapersetter/internal/imagevalidator"
)

func SetWallpaper(input string) error {
	if !imagevalidator.IsImageFile(input) {
		return imagevalidator.ErrNotAnImage
	}
	// Build the qdbus command in chunks to avoid long lines
	qdbusCommand := buildQdbusCommand(input)

	// Execute the qdbus command
	cmd := exec.Command("bash", "-c", qdbusCommand)
	err := cmd.Run()

	// Log and return the error if something goes wrong
	if err != nil {
		log.Printf("Error setting wallpaper: %v", err)
		return fmt.Errorf("failed to set wallpaper: %w", err)
	}

	log.Printf("Wallpaper set to: file://%s", input)
	return nil
}

func buildQdbusCommand(input string) string {
	var sb strings.Builder

	// Build the qdbus command using a string builder
	sb.WriteString("qdbus org.kde.plasmashell /PlasmaShell org.kde.PlasmaShell.evaluateScript \"")
	sb.WriteString("var allDesktops = desktops();")
	sb.WriteString("for (i=0;i<allDesktops.length;i++) {")
	sb.WriteString("    d = allDesktops[i];")
	sb.WriteString("    d.wallpaperPlugin = 'org.kde.image';")
	sb.WriteString("    d.currentConfigGroup = Array('Wallpaper', 'org.kde.image', 'General');")
	sb.WriteString(fmt.Sprintf("    d.writeConfig('Image', 'file://%s')", input))
	sb.WriteString("}\"")

	return sb.String()
}
