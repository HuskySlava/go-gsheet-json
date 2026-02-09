package convert

import "fmt"

type Row struct {
	Key   string
	Value interface{}
}

func FlattenJSONToRows(data interface{}) []Row {
	var rows []Row
	flatten(data, "", &rows)
	return rows
}

func flatten(data interface{}, prefix string, rows *[]Row) {
	switch v := data.(type) {
	// Handle objects, recursively
	case map[string]interface{}:
		for k, val := range v {
			newKey := k
			if prefix != "" {
				newKey = prefix + "." + k
			}
			flatten(val, newKey, rows)
		}
	// Handle arrays, recursively
	case []interface{}:
		for i, val := range v {
			newKey := fmt.Sprintf("%s.%d", prefix, i)
			flatten(val, newKey, rows)
		}
	default:
		*rows = append(*rows, Row{Key: prefix, Value: v})
	}
}
