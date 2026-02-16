package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"go-sheet-json/convert"
	"go-sheet-json/gsheet"
	"log"
	"os"
)

var toGSheetCommand = &cobra.Command{
	Use:   "togsheet",
	Short: "Convert json to 2 column Google Sheet (key/s, value)",
	Run:   toGSheet,
}

var fromFileName string
var toSheetID string
var toSheetRange string

func init() {
	rootCmd.AddCommand(toGSheetCommand)
	toGSheetCommand.Flags().StringVarP(&fromFileName, "file", "f", "", ".json input filename")
	_ = toGSheetCommand.MarkFlagRequired("file")

	toGSheetCommand.Flags().StringVarP(&toSheetID, "sheet", "s", "", "Google sheet ID")
	_ = toGSheetCommand.MarkFlagRequired("sheet")

	toGSheetCommand.Flags().StringVarP(&toSheetRange, "range", "r", "", "Google sheet range - The A1 notation https://developers.google.com/workspace/sheets/api/guides/concepts#cell")
	_ = toGSheetCommand.MarkFlagRequired("range")
}

func toGSheet(cmd *cobra.Command, args []string) {
	fileBytes, err := os.ReadFile(fromFileName)
	if err != nil {
		log.Fatalf("Unable to read file %s: %v", fromFileName, err)
	}

	var jsonData any

	err = json.Unmarshal(fileBytes, &jsonData)
	if err != nil {
		log.Fatal("Failed to parse json: ", err)
	}

	flatten := convert.FlattenJSONToRows(jsonData)

	ctx, cancel := context.WithTimeout(context.Background(), cfg.Timeout)
	defer cancel()

	c, err := gsheet.NewClient(ctx, cfg)
	if err != nil {
		log.Fatal("Failed to connect to google sheet: ", err)
	}

	err = c.WriteSheetRows(toSheetID, toSheetRange, convert.RowsToSlices(flatten))
	if err != nil {
		log.Fatal("Failed writing to google sheet: ", err)
	}

	fmt.Println("Done")
}
