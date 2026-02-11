package gheet

import (
	"context"
	"go-sheet-json/config"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type Client struct {
	Service *sheets.Service
}

func NewClient(cfg *config.Config) (*Client, error) {

	ctx, cancel := context.WithTimeout(context.Background(), cfg.Timeout)
	defer cancel()

	srv, err := sheets.NewService(ctx,
		option.WithCredentialsFile("credentials.json"),
		option.WithScopes(sheets.SpreadsheetsScope),
	)

	return &Client{
		Service: srv,
	}, err
}
