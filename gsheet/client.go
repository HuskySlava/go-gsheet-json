package gheet

import (
	"context"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type Client struct {
	Service *sheets.Service
}

func NewClient() (*Client, error) {

	ctx := context.Background()

	srv, err := sheets.NewService(ctx,
		option.WithCredentialsFile("credentials.json"),
		option.WithScopes(sheets.SpreadsheetsScope),
	)

	return &Client{
		Service: srv,
	}, err
}
