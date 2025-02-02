package advise

import "github.com/spf13/cobra"

func Command() *cobra.Command {
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
