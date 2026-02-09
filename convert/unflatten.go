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
		_, ok := v[0].(string)
		if !ok {
			return fmt.Errorf("row %d: invalid key type, expected: string, got %T", i, v[0])
		}
	}
	return nil
}
