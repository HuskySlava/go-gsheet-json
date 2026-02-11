package gsheet

import (
	"context"
	"fmt"
	"go-sheet-json/config"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type Client struct {
	service *sheets.Service
}

func NewClient(ctx context.Context, cfg *config.Config) (*Client, error) {

	srv, err := sheets.NewService(ctx,
		option.WithAuthCredentialsFile(option.ServiceAccount, cfg.ServiceAccountFilePath),
		option.WithScopes(sheets.SpreadsheetsScope),
	)

	if err != nil {
		return nil, err
	}

	return &Client{
		service: srv,
	}, nil
}

func (c *Client) Test() {
	fmt.Println("Hello")
}
