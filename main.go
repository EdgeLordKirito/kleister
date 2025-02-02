package main

import (
	"os"

	"github.com/EdgeLordKirito/wallpapersetter/internal/appinfo"
	"github.com/EdgeLordKirito/wallpapersetter/internal/commands/advise"
	"github.com/EdgeLordKirito/wallpapersetter/internal/commands/paste"
	"github.com/EdgeLordKirito/wallpapersetter/internal/commands/schedule"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use: appinfo.AppName,
	}
	//rootCmd.AddCommand(chartCommand())
	rootCmd.AddCommand(paste.Command())
	rootCmd.AddCommand(schedule.Command())
	rootCmd.AddCommand(advise.Command())

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1) // Let Cobra handle printing the error
	}
}
