package cmd

import "github.com/spf13/cobra"

var toGSheetCommand = &cobra.Command{
	Use:   "togsheet",
	Short: "Convert json to 2 column Google Sheet (key/s, value)",
	Run:   toGSheet,
}

func init() {
	rootCmd.AddCommand(toGSheetCommand)
}

func toGSheet(cmd *cobra.Command, args []string) {

}
