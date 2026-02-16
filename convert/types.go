package convert

type Row struct {
	Key   string
	Value any
}

func (r Row) ToSlice() []any {
	return []any{r.Key, r.Value}
}

func RowsToSlices(rows []Row) [][]any {
	raw := make([][]any, len(rows))
	for i, r := range rows {
		raw[i] = r.ToSlice()
	}
	return raw
}
