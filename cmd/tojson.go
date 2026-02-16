package cmd

import (
	"context"
	"github.com/spf13/cobra"
	"go-sheet-json/convert"
	"go-sheet-json/gsheet"
	"log"
	"os"
	"path/filepath"
)

var toJsonCommand = &cobra.Command{
	Use:   "tojson",
	Short: "Convert 2 column (key/s, value) Google Sheet to json",
	Run:   toJson,
}

var toFileName string
var fromSheetID string
var fromSheetRange string

func init() {

	rootCmd.AddCommand(toJsonCommand)

	toJsonCommand.Flags().StringVarP(&toFileName, "file", "f", "", ".json output filename")
	_ = toJsonCommand.MarkFlagRequired("file")

	toJsonCommand.Flags().StringVarP(&fromSheetID, "sheet", "s", "", "Google sheet ID")
	_ = toJsonCommand.MarkFlagRequired("sheet")

	toJsonCommand.Flags().StringVarP(&fromSheetRange, "range", "r", "", "Google sheet range - The A1 notation https://developers.google.com/workspace/sheets/api/guides/concepts#cell")
	_ = toJsonCommand.MarkFlagRequired("range")
}

func toJson(cmd *cobra.Command, args []string) {
	ctx, cancel := context.WithTimeout(context.Background(), cfg.Timeout)
	defer cancel()

	c, err := gsheet.NewClient(ctx, cfg)
	if err != nil {
		log.Fatal("Failed to connect to google sheet: ", err)
	}

	d, err := c.ReadSheetRows(fromSheetID, fromSheetRange)
	if err != nil {
		log.Fatal("Failed reading from google sheet: ", err)
	}

	r, err := convert.UnflattenRowsToJSON(d)
	if err != nil {
		log.Fatal("Failed flattening data from google sheet: ", err)
	}

	err = os.MkdirAll(filepath.Dir(toFileName), 0755)
	if err != nil {
		log.Fatal("Failed to create folder for file path: ", err)
	}

	err = os.WriteFile(toFileName, r, 0644)
	if err != nil {
		log.Fatal("Failed to write JSON data to file: ", err)
	}
}
