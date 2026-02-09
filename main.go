package main

import (
	"encoding/json"
	"fmt"
	"go-sheet-json/convert"
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

	flatten := convert.FlattenJSONToRows(data)
	fmt.Println("Flatten:", flatten)

	slices := convert.RowsToSlices(flatten)
	unflatten, err := convert.UnflattenRowsToJSON(slices)
	if err != nil {
		log.Fatal("Failed to parse rows")
	}
	fmt.Println("Flatten:", string(unflatten))
}
