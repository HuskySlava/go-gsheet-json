package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func handleJSONObject(o interface{}, prefix string) {
	m, _ := o.(map[string]interface{})
	for k, v := range m {
		switch v.(type) {
		case map[string]interface{}:
			handleJSONObject(v, prefix+"."+k)
		case string:
			fmt.Printf("Key: %s is a STRING: %s\n", prefix+k, v)
		case float64:
			fmt.Printf("Key: %s is a NUMBER: %f\n", prefix+k, v)
		default:
			fmt.Printf("Key: %s is something else (%T)\n", prefix+k, v)
		}
	}
}

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

	handleJSONObject(data, "")

}
