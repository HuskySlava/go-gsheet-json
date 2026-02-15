package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Server mode",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Running serve cmd")
		// TODO..
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
