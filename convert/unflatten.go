package convert

import (
	"encoding/json"
	"fmt"
	"strings"
)

func UnflattenRowsToJSON(rows [][]any) ([]byte, error) {
	err := validateRows(rows)
	if err != nil {
		return nil, fmt.Errorf("unable to parse rows to json: %w", err)
	}

	root := map[string]any{}

	for _, v := range rows {
		key := v[0].(string)
		value := v[1]

		parts := strings.Split(key, ".")
		current := root

		// Check "part" within the current map - if it's a nested map.
		// If it doesn't exist or isn't a map, create it
		for _, part := range parts[:len(parts)-1] {
			next, exists := current[part].(map[string]any)
			if !exists {
				next = map[string]any{}
				current[part] = next
			}
			current = next
		}

		current[parts[len(parts)-1]] = value
	}

	return json.MarshalIndent(root, "", "  ")
}

func validateRows(rows [][]any) error {
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
