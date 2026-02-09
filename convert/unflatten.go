package convert

import "fmt"

func UnflattenRowsToJSON(rows [][]interface{}) ([]byte, error) {
	err := validateRows(rows)
	if err != nil {
		return nil, fmt.Errorf("unable to parse rows to json: %w", err)
	}

	return nil, nil
}

func validateRows(rows [][]interface{}) error {
	for i, v := range rows {
		if len(v) != 2 {
			return fmt.Errorf("row %d: invalid number of columns, expected: 2, got %d", i, len(v))
		}
	}
	return nil
}
