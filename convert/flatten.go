package convert

import "fmt"

func FlattenJSONToRows(data any) []Row {
	var rows []Row
	flatten(data, "", &rows)
	return rows
}

func flatten(data any, prefix string, rows *[]Row) {
	switch v := data.(type) {
	// Handle objects, recursively
	case map[string]any:
		for k, val := range v {
			newKey := k
			if prefix != "" {
				newKey = prefix + "." + k
			}
			flatten(val, newKey, rows)
		}
	// Handle arrays, recursively
	case []any:
		for i, val := range v {
			newKey := fmt.Sprintf("%s.%d", prefix, i)
			flatten(val, newKey, rows)
		}
	default:
		*rows = append(*rows, Row{Key: prefix, Value: v})
	}
}
