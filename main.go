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

func updateGSheetWithJSON(json interface{}, prefix string, csv *[][]interface{}) *[][]interface{} {

	m, _ := json.(map[string]interface{})

	for k, v := range m {
		switch v.(type) {
		case map[string]interface{}:
			updateGSheetWithJSON(v, prefix+k+".", csv)
		default:
			row := []interface{}{
				prefix + k,
				v,
			}
			*csv = append(*csv, row)
		}
	}

	return csv
}
