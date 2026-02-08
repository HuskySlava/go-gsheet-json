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
	case map[string]interface{}:
		for k, val := range v {
			newKey := k
			if prefix != "" {
				newKey = prefix + "." + k
			}
			updateGSheetWithJSON(val, newKey, gsheet)
		}
	case []interface{}:
		for i, val := range v {
			newKey := fmt.Sprintf("%s.%d", prefix, i)
			updateGSheetWithJSON(val, newKey, gsheet)
		}
	default:
		row := []interface{}{prefix, v}
		*gsheet = append(*gsheet, row)
	}
}
