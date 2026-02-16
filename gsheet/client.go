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

func (c *Client) ReadSheetRows(spreadsheetID string, sheetRange string) ([][]any, error) {
	res, err := c.service.Spreadsheets.Values.Get(spreadsheetID, sheetRange).Do()
	if err != nil {
		return nil, fmt.Errorf("unable to read spreadsheet: %w", err)
	}
	return res.Values, nil
}

func (c *Client) WriteSheetRows(spreadsheetID string, sheetRange string, values [][]any) error {
	vr := &sheets.ValueRange{
		Values: values,
	}
	_, err := c.service.Spreadsheets.Values.Update(spreadsheetID, sheetRange, vr).
		ValueInputOption("RAW").
		Do()
	if err != nil {
		return fmt.Errorf("unable to update spreadsheet: %w", err)
	}
	return nil
}
