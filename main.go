package main

import (
	"context"
	"go-sheet-json/config"
	"go-sheet-json/gsheet"
	"log"
)

func main() {
	cfg, err := config.Load("config.yaml")
	if err != nil {
		log.Fatal("Failed to load config: ", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), cfg.Timeout)
	defer cancel()

	gsheetClient, err := gsheet.NewClient(ctx, cfg)
	if err != nil {
		log.Fatal("Failed to initialize google sheet client:", err)
	}

	gsheetClient.Test()
}
