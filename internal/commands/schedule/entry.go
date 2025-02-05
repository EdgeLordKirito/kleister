package schedule

import "github.com/spf13/cobra"

var (
	path     string
	interval string
)

const (
	timeFormatFlagName string = "chronoformat"
)

func Command() *cobra.Command {
	var scheduleCMD = cobra.Command{
		Use: "schedule",
		//TODO: figure ot the short and long description
		Short: "",
		Long:  "",
		//TODO: configure args and set the function to run
		RunE: Run,
	}

	scheduleCMD.Flags().StringVarP(&path, "input",
		"i", "", "file path or directory path")
	scheduleCMD.Flags().StringVarP(&interval, timeFormatFlagName,
		"f", "", "time format string")
	scheduleCMD.MarkFlagRequired(timeFormatFlagName)

	return &scheduleCMD
}
