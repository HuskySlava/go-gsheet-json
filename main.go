package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.ReadFile("./test.json")
	if err != nil {
		fmt.Println("error", err)
	}

	var data interface{}

	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatal("failed to parse json", err)
	}

	var gsheet [][]interface{}
	updateGSheetWithJSON(data, "", &gsheet)

	fmt.Println(gsheet)
}

func updateGSheetWithJSON(data interface{}, prefix string, gsheet *[][]interface{}) {
	switch v := data.(type) {
	// Handle objects, recursively
	case map[string]interface{}:
		for k, val := range v {
			newKey := k
			if prefix != "" {
				newKey = prefix + "." + k
			}
			updateGSheetWithJSON(val, newKey, gsheet)
		}
	// Handle arrays, recursively
	case []interface{}:
		for i, val := range v {
			newKey := fmt.Sprintf("%s.%d", prefix, i)
			updateGSheetWithJSON(val, newKey, gsheet)
		}
	default:
		// Append key, value to a row
		row := []interface{}{prefix, v}
		// Append row to a slice of rows
		*gsheet = append(*gsheet, row)
	}
}
