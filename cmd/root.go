package cmd

import (
	"github.com/spf13/cobra"
	"go-sheet-json/config"
	"log"
	"os"
)

var cfg *config.Config

var rootCmd = &cobra.Command{
	Use:   "go-sheet-json",
	Short: "My app stuff",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		var err error
		cfg, err = config.Load("config.yaml")
		if err != nil {
			log.Fatal(err)
		}
		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
