package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var toJsonCommand = &cobra.Command{
	Use:   "tojson",
	Short: "Convert 2 column (key/s, value) Google Sheet to json",
	Run:   toJson,
}

func init() {
	rootCmd.AddCommand(toJsonCommand)
}

func toJson(cmd *cobra.Command, args []string) {
	fmt.Println("To json...", cmd)
}
