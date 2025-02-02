package main

import (
	"os"

	"github.com/EdgeLordKirito/wallpapersetter/internal/appinfo"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use: appinfo.AppName,
	}
	//rootCmd.AddCommand(chartCommand())
	rootCmd.AddCommand(pasteCommand())
	rootCmd.AddCommand(scheduleCommand())
	rootCmd.AddCommand(adviseCommand())

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1) // Let Cobra handle printing the error
	}
}

// Sets wallpaper once
func pasteCommand() *cobra.Command {
	var tuiCmd = &cobra.Command{
		Use: "paste",
		//TODO: figure ot the short and long description
		Short: "Command for setting the Wallpaper",
		Long:  "paste sets the Wallpaper using the specified path",
		//TODO: configure args and set the function to run
		Args: cobra.ExactArgs(1),
		RunE: nil,
	}

	return tuiCmd
}

// schedules wallpaper changes
func scheduleCommand() *cobra.Command {
	var scheduleCMD = cobra.Command{
		Use: "schedule",
		//TODO: figure ot the short and long description
		Short: "",
		Long:  "",
		//TODO: configure args and set the function to run
		Args: nil,
		RunE: nil,
	}
	return &scheduleCMD
}

// opens the wui
func adviseCommand() *cobra.Command {
	var adviseCMD = cobra.Command{
		Use: "advise",
		//TODO: figure ot the short and long description
		Short: "",
		Long:  "",
		//TODO: configure args and set the function to run
		Args: nil,
		RunE: nil,
	}
	return &adviseCMD
}
