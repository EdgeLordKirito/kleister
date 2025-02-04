package paste

import "github.com/spf13/cobra"

var (
	path     string
	ordering string
)

func Command() *cobra.Command {
	var pasteCmd = &cobra.Command{
		Use: "paste",
		//TODO: figure ot the short and long description
		Short: "Command for setting the Wallpaper",
		Long:  "paste sets the Wallpaper using the specified path",
		//TODO: configure args and set the function to run
		RunE: Run,
	}

	pasteCmd.Flags().StringVarP(&path, "input",
		"i", "", "file path or directory path")
	pasteCmd.Flags().StringVarP(&ordering, "output",
		"o", "random", "Ordering mode")

	return pasteCmd
}
