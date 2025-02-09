package stop

import "github.com/spf13/cobra"

func Command() *cobra.Command {
	var stopCMD = cobra.Command{
		Use: "stop",
		//TODO: figure ot the short and long description
		Short: "",
		Long:  "",
		//TODO: configure args and set the function to run
		RunE: Run,
	}

	return &stopCMD
}
