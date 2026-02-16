package convert

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
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

	converted := mapsToSlices(root)

	return json.MarshalIndent(converted, "", "  ")
}

// mapsToSlices recursively walks the tree and converts maps whose keys
// are all consecutive integers starting from 0 into slices.
func mapsToSlices(data any) any {
	m, ok := data.(map[string]any)
	if !ok {
		return data
	}

	// Recurse first so children are converted before we inspect this level
	for k, v := range m {
		m[k] = mapsToSlices(v)
	}

	if !isSequentialIntKeys(m) {
		return m
	}

	slice := make([]any, len(m))
	for k, v := range m {
		i, _ := strconv.Atoi(k)
		slice[i] = v
	}
	return slice
}

// isSequentialIntKeys returns true if all keys are integers 0..len-1.
func isSequentialIntKeys(m map[string]any) bool {
	if len(m) == 0 {
		return false
	}

	keys := make([]int, 0, len(m))
	for k := range m {
		n, err := strconv.Atoi(k)
		if err != nil {
			return false
		}
		keys = append(keys, n)
	}

	sort.Ints(keys)
	for i, k := range keys {
		if k != i {
			return false
		}
	}
	return true
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
